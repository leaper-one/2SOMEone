// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	accounts "github.com/leaper-one/2SOMEone/api/user-api/internal/handler/accounts"
	app "github.com/leaper-one/2SOMEone/api/user-api/internal/handler/app"
	user "github.com/leaper-one/2SOMEone/api/user-api/internal/handler/user"
	"github.com/leaper-one/2SOMEone/api/user-api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/sent-phone-code/:phone",
				Handler: user.SentPhoneCodeHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/signup",
				Handler: user.SignUpHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/signin",
				Handler: user.SignInHandler(serverCtx),
			},
		},
		rest.WithPrefix("/user/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/me",
				Handler: accounts.GetMeHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/setInfo",
				Handler: accounts.SetInfoHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/account/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/get-userid-by-buid",
				Handler: app.GetUserIdByBuidHandler(serverCtx),
			},
		},
		rest.WithPrefix("/app/v1"),
	)
}
