package service

import (
	"2SOMEone/core"
	"2SOMEone/store/msg"
	"2SOMEone/util"
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"

	dysmsapi "github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
)

func NewMsgService(db *util.DB) *MsgService {
	return &MsgService{
		db: db,
	}
}

type MsgService struct {
	db *util.DB
}

// Send random code to phone via aliyun msg api
func (a *MsgService) SendPhoneCode(ctx context.Context, phone string) (string, uint, error) {
	// 初始化 msg store
	msgStore := msg.New(a.db)
	// 初始化阿里云短信服务客户端
	client, err := dysmsapi.NewClientWithAccessKey("cn-hangzhou", "LTAI5tREMX8wtEQoaSgGki4Z", "YGtpz8dZWTrWQqDm4fk4NlsaFHJNCW")
	if err != nil {
		return "", 0, err
	}

	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"

	// 生成六位随机码
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vcode := fmt.Sprintf("%06v", rnd.Int31n(1000000)) // store

	request.PhoneNumbers = phone
	request.SignName = "ABC商城"
	request.TemplateCode = "SMS_205575254"
	request.TemplateParam = "{\"code\":" + "\"" + vcode + "\"}"

	respon, err := client.SendSms(request)
	if err != nil {
		return "", 0, err
	}
	if respon.Code != "OK" {
		return "", 0, errors.New(respon.Message)
	}

	msg := &core.Message{Type: 0, Phone: phone, Content: "", Code: vcode}
	err = msgStore.Create(ctx, msg)
	if err != nil {
		return "", 0, err
	}

	return msg.Code, msg.ID, nil
}

// CheckPhoneCode check random code with phone and msg_id
func (a *MsgService) CheckPhoneCode(ctx context.Context, phone, code string, msg_id uint) (bool, error) {
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
