package message

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/leaper-one/2SOMEone/store/msg"
	"github.com/leaper-one/2SOMEone/util"

	dysmsapi "github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
)

var (
	config = util.LoadConfig("./config.yaml", &Config{}).(*Config)
	dbc    = util.OpenDB("./message.db")
)

type MessageService struct{}

// config struct
// app, alimsg, grpcset
type Config struct {
	AliMsg struct {
		RegionId        string `yaml:"region_id"`
		AccessKeyId     string `yaml:"access_key_id"`
		AccessKeySecret string `yaml:"access_key_secret"`
		SignName        string `yaml:"sign_name"`
		TemplateCode    string `yaml:"template_code"`
	}
}

// params: db, region_id, access_key_id, access_key_secret
func NewMsgService(db *util.DB, msg_config ...string) *MsgService {
	// 初始化阿里云短信服务客户端
	if len(msg_config) > 0 {
		client, err := dysmsapi.NewClientWithAccessKey(msg_config[0], msg_config[1], msg_config[2])
		if err != nil {
			panic(err)
		}
		return &MsgService{
			client: client,
			db:     db,
		}
	} else {
		return &MsgService{
			db: db,
		}
	}
}

type MsgService struct {
	db     *util.DB
	client *dysmsapi.Client
}

func SendPhoneCode(ctx context.Context, phone string) (string, error) {
	msgService := NewMsgService(dbc, config.AliMsg.RegionId, config.AliMsg.AccessKeyId, config.AliMsg.AccessKeySecret)
	return msgService.sendPhoneCode(ctx, phone)
}

// Send random code to phone via aliyun msg api
func (a *MsgService) sendPhoneCode(ctx context.Context, phone string) (string, error) {

	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"

	// 生成六位随机码
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vcode := fmt.Sprintf("%06v", rnd.Int31n(1000000)) // store

	request.PhoneNumbers = phone
	request.SignName = config.AliMsg.SignName
	request.TemplateCode = config.AliMsg.TemplateCode
	request.TemplateParam = "{\"code\":" + "\"" + vcode + "\"}"

	response, err := a.client.SendSms(request)
	if err != nil {
		return "", err
	}
	if response.Code != "OK" {
		return "", errors.New(response.Message)
	}

	return vcode, nil
}

func CheckPhoneCode(ctx context.Context, phone, code string, msg_id uint) (bool, error) {
	msgService := NewMsgService(dbc, config.AliMsg.RegionId, config.AliMsg.AccessKeyId, config.AliMsg.AccessKeySecret)
	return msgService.checkPhoneCode(ctx, phone, code, msg_id)
}

// CheckPhoneCode check random code with phone and msg_id
func (a *MsgService) checkPhoneCode(ctx context.Context, phone, code string, msg_id uint) (bool, error) {
	msgStore := msg.New(a.db)
	msg, err := msgStore.Find(ctx, msg_id, phone)
	if err != nil {
		return false, err
	}
	if msg == nil && err == nil {
		return false, errors.New("无此短信 ID")
	}
	if msg.Code == code {
		return true, nil
	} else {
		return false, errors.New("验证码错误")
	}
}
