package user

import (
	"net/http"

	"github.com/leaper-one/2SOMEone/api/user-api/internal/logic/user"
	"github.com/leaper-one/2SOMEone/api/user-api/internal/svc"
	"github.com/leaper-one/2SOMEone/api/user-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"

)

func SentPhoneCodeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SentPhoneCodeReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewSentPhoneCodeLogic(r.Context(), svcCtx)
		resp, err := l.SentPhoneCode(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
