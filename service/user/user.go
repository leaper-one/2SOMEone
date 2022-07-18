package account

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/leaper-one/2SOMEone/core"
	"github.com/leaper-one/2SOMEone/store/bili_user"
	"github.com/leaper-one/2SOMEone/store/user"
	"github.com/leaper-one/2SOMEone/util"
	// "github.com/leaper-one/2SOMEone/service/message"

	"github.com/gofrs/uuid"
)

var (
	dbc    = util.OpenDB("./user.db")
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

func SignUpByPhone(ctx context.Context, phone, password string) error {
	userService := NewUserService(dbc)
	err := userService.signUpByPhone(ctx, phone, password)
	if err != nil {
		return errors.New("注册失败")
	}
	return nil
}

func (a *UserService) signUpByPhone(ctx context.Context, phone, password string) error {
	// TODO: 检车用户是否存在

	// 创建用户
	userStore := user.NewUserStore(a.db)
	user_id, _ := uuid.NewV1()

	user := &core.BasicUser{UserID: user_id.String(), Phone: phone, Name: phone}

	// 密码加密
	user.Password, _ = util.HashPassword(password)

	err := userStore.Save(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

// func (a *UserService) Login(ctx context.Context, token string) (*core.BasicUser, error) {}
func (a *UserService) Auth(ctx context.Context, phone, password string) (string, error) {
	userStore := user.NewUserStore(a.db)
	user, err := userStore.FindByPhone(ctx, phone)
	if err != nil {
		return "", err
	} else if user == nil && err == nil {
		return "", errors.New("该手机号未绑定")
	}
	if util.CheckPasswordHash(password, user.Password) {
		token, err := util.GenerateToken(user.UserID, user.Phone, time.Hour*7*24)
		if err != nil {
			return "", err
		}
		return token, nil
	}
	return "", errors.New("密码错误")
}

func (a *UserService) SetInfo(ctx context.Context, user_id, name, avatar, buid string) error {
	userStore := user.NewUserStore(a.db)
	user, err := userStore.FindByUserID(ctx, user_id)
	if err != nil {
		return err
	} else if user == nil && err == nil {
		return errors.New("无此用户")
	}

	if name != "" {
		user.Name = name
	}

	if avatar != "" {
		user.Avatar = avatar
	}

	if buid != "" {
		params := url.Values{}
		Url, err := url.Parse("https://api.bilibili.com/x/space/acc/info")
		if err != nil {
			return err
		}
		params.Set("mid", buid) //如果参数中有中文参数,这个方法会进行URLEncode
		Url.RawQuery = params.Encode()
		urlPath := Url.String()
		resp, err := http.Get(urlPath)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		var bli_user_info core.BiliUserInfo
		err = json.Unmarshal(body, &bli_user_info)
		if err != nil {
			return err
		}

		// 写入 BiliUser
		biliUserStore := bili_user.NewBiliUserStore(a.db)
		biliUserStore.Save(ctx, &core.BiliUser{
			UserID:      user_id,
			Buid:        int64(bli_user_info.Data.Mid),
			LiveRoomID:  int64(bli_user_info.Data.LiveRoom.RoomID),
			LiveRoomUrl: bli_user_info.Data.LiveRoom.Url,
		})

		if name == "" {
			user.Name = bli_user_info.Data.Name
		}

		if avatar == "" {
			user.Avatar = bli_user_info.Data.Face
		}
	}

	err = userStore.Save(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (a *UserService) GetMe(ctx context.Context, user_id string) (*core.BasicUser, error) {
	userStore := user.NewUserStore(a.db)
	// user, err := userStore.FindByPhone(ctx, .Phone)
	user, err := userStore.FindByUserID(ctx, user_id)
	if err != nil {
		return nil, err
	} else if user == nil && err == nil {
		return nil, errors.New("无此用户")
	}
	return user, nil
}

func (a *UserService) FindByBuid(ctx context.Context, buid int64) (*core.BiliUser, error) {
	biliUserStore := bili_user.NewBiliUserStore(a.db)
	buser, err := biliUserStore.FindByBuid(ctx, buid)
	if err != nil {
		return nil, err
	} else if buser == nil && err == nil {
		return nil, errors.New("无此用户")
	}
	return buser, nil
}