package login

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

type GetSecurityTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

type GetSecurityTokenYxyResp struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Data       struct {
		Level         uint8  `json:"level"`
		SecurityToken string `json:"securityToken"`
	} `json:"data"`
	Success bool `json:"success"`
}

func NewGetSecurityTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSecurityTokenLogic {
	return &GetSecurityTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSecurityTokenLogic) GetSecurityToken(req *types.GetSecurityTokenReq) (resp *types.GetSecurityTokenResp, err error) {
	yxyReq, yxyHeaders := yxyClient.GetYxyBaseReqParam(req.DeviceID)
	yxyReq["sceneCode"] = "app_user_login"

	var yxyResp GetSecurityTokenYxyResp
	r, err := yxyClient.HttpSendPost(consts.GET_SECURITY_TOKEN_URL, yxyReq, yxyHeaders, &yxyResp)
	if err != nil {
		return nil, err
	}

	if yxyResp.StatusCode != 0 {
		return nil, xerr.WithCode(xerr.ErrUnknown, fmt.Sprintf("yxy response: %v", r))
	}

	return &types.GetSecurityTokenResp{
		Level:         yxyResp.Data.Level,
		SecurityToken: yxyResp.Data.SecurityToken,
	}, nil
}
