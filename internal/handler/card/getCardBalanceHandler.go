package card

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"yxy-go/internal/logic/card"
	"yxy-go/internal/svc"
	"yxy-go/internal/types"
	"yxy-go/pkg/response"
)

func GetCardBalanceHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetCardBalanceReq
		if err := httpx.Parse(r, &req); err != nil {
			response.ParamErrorResponse(r, w, err)
			return
		}

		l := card.NewGetCardBalanceLogic(r.Context(), svcCtx)
		resp, err := l.GetCardBalance(&req)
		response.HttpResponse(r, w, resp, err)
	}
}
