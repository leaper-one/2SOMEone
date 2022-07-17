package accounts

import (
	"net/http"

	"github.com/leaper-one/2SOMEone/api/user-api/internal/logic/accounts"
	"github.com/leaper-one/2SOMEone/api/user-api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SentPhoneCodeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := accounts.NewSentPhoneCodeLogic(r.Context(), svcCtx)
		resp, err := l.SentPhoneCode()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
