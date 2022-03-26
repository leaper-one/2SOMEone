package service

import (
	"2SOMEone/core"
	"2SOMEone/store/user"
	"2SOMEone/util"
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"

	dysmsapi "github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/gofrs/uuid"
)

func NewUserService(
	db *util.DB,
) *UserService {
	return &UserService{
		db: db,
	}
}

type UserService struct {
	db *util.DB
}

func (a *UserService) SendPhoneCode(ctx context.Context, phone string) (string, error) {
	client, err := dysmsapi.NewClientWithAccessKey("cn-hangzhou", "LTAI5tREMX8wtEQoaSgGki4Z", "YGtpz8dZWTrWQqDm4fk4NlsaFHJNCW")
	if err != nil {
		return "", err
	}
	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vcode := fmt.Sprintf("%06v", rnd.Int31n(1000000)) // store

	userStore := user.New(a.db)
	go userStore.Save(ctx, &core.User{Phone: phone, Code: vcode, Role: "informal"})

	request.PhoneNumbers = phone
	request.SignName = "ABC商城"
	request.TemplateCode = "SMS_205575254"
	request.TemplateParam = "{\"code\":" + "\"" + vcode + "\"}"

	respon, err := client.SendSms(request)
	if err != nil {
		return "", err
	}
	fmt.Printf("respon: %+v\n", respon)

	return vcode, nil
}

func (a *UserService) SignUpByPhone(ctx context.Context, sign_user *core.SignUpUser) (*core.User, error) {
	userStore := user.New(a.db)
	user, err := userStore.FindByPhone(ctx, sign_user.Phone)
	if err != nil {
		return nil, err
	} else if user == nil && err == nil {
		return nil, errors.New("无此非正式用户")
	}
	// if code == user.Code {
	if sign_user.Code == "000000" || sign_user.Code == user.Code { // TODO: 测试用
		user_id, _ := uuid.NewV1()
		// user := &core.User{Email: email, Role: "formal", UserID: user_id.String(),Code: ""}
		user.Password, _ = util.HashPassword(sign_user.Password)
		// user.Role = "formal"
		user.Name = user.Phone
		user.UserID = user_id.String()
		user.Code = ""
		err := userStore.Save(ctx, user)
		if err != nil {
			return nil, err
		}
		user.Role = "formal"
		err = userStore.Save(ctx, user)
		if err != nil {
			return nil, err
		}
		return user, nil
	} else {
		return nil, errors.New("验证码错误")
	}
}

// func (a *UserService) Login(ctx context.Context, token string) (*core.User, error) {}
func (a *UserService) Auth(ctx context.Context, login_user *core.LoginUser) (string, error) {
	userStore := user.New(a.db)
	user, err := userStore.FindByPhone(ctx, login_user.Phone)
	if err != nil {
		return "", err
	} else if user == nil && err == nil {
		return "", errors.New("该手机号未绑定")
	}
	if util.CheckPasswordHash(login_user.Password, user.Password) {
		token, err := util.GenerateToken(user.UserID, time.Hour*7*24)
		if err != nil {
			return "", err
		}
		return token, nil
	}
	return "", errors.New("密码错误")
}

func (a *UserService) GetMe(ctx context.Context, user_id string) (*core.UserForMe, error) {
	userStore := user.New(a.db)
	// user, err := userStore.FindByPhone(ctx, .Phone)
	user, err := userStore.FindByUserIDForMe(ctx, user_id)
	if err != nil {
		return nil, err
	} else if user == nil && err == nil {
		return nil, errors.New("无此用户")
	}
	return user, nil
}

func (a *UserService) VisitUser(ctx context.Context, user_name string) (*core.UserForShow, error) {
	userStore := user.New(a.db)
	user, err := userStore.FindByName(ctx, user_name)
	if err != nil {
		return nil, err
	} else if user == nil && err == nil {
		return nil, errors.New("无此用户")
	}
	user_for_show, err := userStore.FindByUserIDForShow(ctx, user.UserID)
	if err != nil {
		return nil, err
	}
	return user_for_show, nil
}
