package logic

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/leaper-one/2SOMEone/core"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/leaper-one/2SOMEone/rpc/user-rpc/internal/svc"
	"github.com/leaper-one/2SOMEone/rpc/user-rpc/types/user"
	"github.com/zeromicro/go-zero/core/logx"
)

type SetInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetInfoLogic {
	return &SetInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  jwt needed in metadata
func (l *SetInfoLogic) SetInfo(in *user.SetInfoRequest) (*user.SetInfoResponse, error) {
	// logx.Info("rpc user-rpc userId: %d", l.ctx.Value("user_id"))
	// user_id := l.ctx.Value("user_id").(string)
	//err := user_service.SetInfo(l.ctx, in.UserId, in.Name, in.Avatar, in.Buid)

	// 根据 user_id 查询 用户数据
	userData, err := l.svcCtx.BasicUsersModel.FindOneByUserId(l.ctx, in.UserId)

	// 根据返回值判断查询是否出错或用户是否存在
	if err != nil {
		return &user.SetInfoResponse{
			Code: 500,
			Msg:  "Fatal",
		}, nil
	}
	if userData == nil {
		return &user.SetInfoResponse{
			Code: 400,
			Msg:  "该用户不存在",
		}, nil
	}

	// 更新用户数据
	if in.Name != "" {
		userData.Name = sql.NullString{
			String: in.Name,
			Valid:  true,
		}
	}

	if in.Avatar != "" {
		userData.Avatar = sql.NullString{
			String: in.Avatar,
			Valid:  true,
		}
	}

	if in.Buid != "" {
		params := url.Values{}
		Url, err := url.Parse("https://api.bilibili.com/x/space/acc/info")
		if err != nil {
			return &user.SetInfoResponse{
				Code: 500,
				Msg:  "Url Parse Error",
			}, nil
		}
		params.Set("mid", in.Buid) // 如果参数中有中文参数,这个方法会进行URLEncode
		Url.RawQuery = params.Encode()
		urlPath := Url.String()
		resp, err := http.Get(urlPath)
		if err != nil {
			return &user.SetInfoResponse{
				Code: 500,
				Msg:  "Http Get Error",
			}, nil
		}
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {

			}
		}(resp.Body)
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return &user.SetInfoResponse{
				Code: 500,
				Msg:  "Read Body Error",
			}, nil
		}
		var bili_user_info core.BiliUserInfo
		err = json.Unmarshal(body, &bili_user_info)
		if err != nil {
			return &user.SetInfoResponse{
				Code: 500,
				Msg:  "Unmarshal Error",
			}, nil
		}

		// 写入 BiliUser 表
		_, err = l.svcCtx.BiliUsersModel.InsertBiliUser(l.ctx, in.UserId, int64(bili_user_info.Data.Mid), int64(bili_user_info.Data.LiveRoom.RoomID), bili_user_info.Data.LiveRoom.Url)
		if err != nil {
			return &user.SetInfoResponse{
				Code: 500,
				Msg:  "Error Insert to BiliUser",
			}, nil
		}

		if in.Name == "" {
			userData.Name = sql.NullString{
				String: bili_user_info.Data.Name,
				Valid:  true,
			}

		}
		if in.Avatar == "" {
			userData.Avatar = sql.NullString{
				String: bili_user_info.Data.Face,
				Valid:  true,
			}
		}
	}

	_, err = l.svcCtx.BasicUsersModel.Insert(l.ctx, userData)

	if err != nil {
		return &user.SetInfoResponse{
			Code: 400,
			Msg:  "set info failed",
		}, err
	}
	return &user.SetInfoResponse{
		Code: 200,
		Msg:  "success",
	}, nil
}
