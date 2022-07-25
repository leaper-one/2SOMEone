package app

import (
	"net/http"

	"github.com/leaper-one/2SOMEone/api/user-api/internal/logic/app"
	"github.com/leaper-one/2SOMEone/api/user-api/internal/svc"
	"github.com/leaper-one/2SOMEone/api/user-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"

)

func GetUserIdByBuidHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetUserIdByBuidReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := app.NewGetUserIdByBuidLogic(r.Context(), svcCtx)
		resp, err := l.GetUserIdByBuid(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
