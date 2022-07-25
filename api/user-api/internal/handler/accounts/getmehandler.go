package accounts

import (
	"net/http"

	"github.com/leaper-one/2SOMEone/api/user-api/internal/logic/accounts"
	"github.com/leaper-one/2SOMEone/api/user-api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetMeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := accounts.NewGetMeLogic(r.Context(), svcCtx)
		resp, err := l.GetMe()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
