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

type GetCaptchaImageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

type GetCaptchaImageYxyResp struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Data       string `json:"data"`
	Success    bool   `json:"success"`
}

func NewGetCaptchaImageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCaptchaImageLogic {
	return &GetCaptchaImageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCaptchaImageLogic) GetCaptchaImage(req *types.GetCaptchaImageReq) (resp *types.GetCaptchaImageResp, err error) {
	yxyReq, yxyHeaders := yxyClient.GetYxyBaseReqParam(req.DeviceID)
	yxyReq["securityToken"] = req.SecurityToken

	var yxyResp GetCaptchaImageYxyResp
	r, err := yxyClient.HttpSendPost(consts.GET_CAPTCHA_IMAGE_URL, yxyReq, yxyHeaders, &yxyResp)
	if err != nil {
		return nil, err
	}

	if yxyResp.StatusCode != 0 {
		errCode := xerr.ErrUnknown
		if yxyResp.Message == "token无效" {
			errCode = xerr.ErrTokenInvalid
		}
		return nil, xerr.WithCode(errCode, fmt.Sprintf("yxy response: %v", r))
	}

	return &types.GetCaptchaImageResp{
		Img: yxyResp.Data,
	}, nil
}
