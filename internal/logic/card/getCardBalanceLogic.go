package card

import (
	"context"
	"fmt"

	"yxy-go/internal/consts"
	"yxy-go/internal/svc"
	"yxy-go/internal/types"
	"yxy-go/internal/utils/yxyClient"
	"yxy-go/pkg/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCardBalanceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCardBalanceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCardBalanceLogic {
	return &GetCardBalanceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type GetCardBalanceYxyResp struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	BizCode    string `json:"bizCode"`
	Success    bool   `json:"success"`
	Data       string `json:"data"`
}

func (l *GetCardBalanceLogic) GetCardBalance(req *types.GetCardBalanceReq) (resp *types.GetCardBalanceResp, err error) {
	yxyReq, yxyHeaders := yxyClient.GetYxyBaseReqParam(req.DeviceID)
	yxyReq["ymId"] = req.UID
	yxyReq["schoolCode"] = consts.SCHOOL_CODE
	yxyReq["walletNo"] = "1"

	var yxyResp GetCardBalanceYxyResp
	r, err := yxyClient.HttpSendPost(consts.GET_CARD_BALANCE_URL, yxyReq, yxyHeaders, &yxyResp)
	if err != nil {
		return nil, err
	}

	if yxyResp.StatusCode != 0 {
		bizCode := yxyResp.BizCode
		errCode := xerr.ErrUnknown
		switch bizCode {
		case "10010":
			errCode = xerr.ErrUserNotFound
		case "10011":
			errCode = xerr.ErrAccountLoggedOut
		case "-1":
			errCode = xerr.ErrNotBindCard
		}
		return nil, xerr.WithCode(errCode, fmt.Sprintf("yxy response: %v", r))
	}

	return &types.GetCardBalanceResp{
		Balance: yxyResp.Data,
	}, nil
}
