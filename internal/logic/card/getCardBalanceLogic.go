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
	Data       string `json:"data"`
	Success    bool   `json:"success"`
}

func (l *GetCardBalanceLogic) GetCardBalance(req *types.GetCardBalanceReq) (resp *types.GetCardBalanceResp, err error) {
	yxyReq, yxyHeaders := yxyClient.GetYxyBaseReqParam(req.DeviceID)
	yxyReq["ymId"] = req.UID
	yxyReq["schoolCode"] = consts.SCHOOL_CODE
	yxyReq["walletNo"] = "1"

	yxyReq["token"] = req.Token
	if req.Token == "" {
		yxyReq["token"] = yxyClient.GenRanmonFakeMd5Token()
	}

	var yxyResp GetCardBalanceYxyResp
	r, err := yxyClient.HttpSendPost(consts.GET_CARD_BALANCE_URL, yxyReq, yxyHeaders, &yxyResp)
	if err != nil {
		return nil, err
	}

	if yxyResp.StatusCode != 0 {
		errCode := xerr.ErrUnknown
		if yxyResp.Message == "登录已过期，请重新登录[user no find]" {
			errCode = xerr.ErrUserNotFound
		} else if yxyResp.Message == "您的账号已被登出，请重新登录[deviceId changed]" {
			errCode = xerr.ErrAccountLoggedOut
		}
		return nil, xerr.WithCode(errCode, fmt.Sprintf("yxy response: %v", r))
	}

	return &types.GetCardBalanceResp{
		Balance: yxyResp.Data,
	}, nil
}
