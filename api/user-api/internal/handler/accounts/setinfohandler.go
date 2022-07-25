package accounts

import (
	"net/http"

	"github.com/leaper-one/2SOMEone/api/user-api/internal/logic/accounts"
	"github.com/leaper-one/2SOMEone/api/user-api/internal/svc"
	"github.com/leaper-one/2SOMEone/api/user-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SetInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SetInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := accounts.NewSetInfoLogic(r.Context(), svcCtx)
		resp, err := l.SetInfo(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
