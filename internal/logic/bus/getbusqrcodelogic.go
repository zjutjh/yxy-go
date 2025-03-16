package bus

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"yxy-go/internal/consts"
	"yxy-go/internal/svc"
	"yxy-go/internal/types"
	"yxy-go/internal/utils/yxyClient"
	"yxy-go/pkg/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetBusQrcodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetBusQrcodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBusQrcodeLogic {
	return &GetBusQrcodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type GetBusQrcodeYxyResp struct {
	Qrcode string `json:"qrcode"`
	Detail struct {
		Code string `json:"code"`
		Msg  string `json:"msg"`
	} `json:"detail"`
}

func (l *GetBusQrcodeLogic) GetBusQrcode(req *types.GetBusQrcodeReq) (resp *types.GetBusQrcodeResp, err error) {
	var yxyResp GetBusQrcodeYxyResp
	client := yxyClient.GetClient()
	r, err := client.R().
		SetHeader("Authorization", req.Token).
		Get(consts.GET_BUS_QRCODE_URL)

	if err != nil {
		log.Printf("Error sending request to %s: %v\n", consts.GET_BUS_RECORD_URL, err)
		return nil, xerr.WithCode(xerr.ErrHttpClient, err.Error())
	}

	err = json.Unmarshal(r.Body(), &yxyResp)
	if err != nil {
		log.Fatalf("Error unmarshaling JSON: %v", err)
		return nil, xerr.WithCode(xerr.ErrHttpClient, err.Error())
	}

	if r.StatusCode() == 400 {
		if yxyResp.Detail.Code == "AUTH_FAIL" {
			return nil, xerr.WithCode(xerr.ErrTokenInvalid, "权限验证失败")
		} else {
			return nil, xerr.WithCode(xerr.ErrHttpClient, fmt.Sprintf("yxy response: %v", r))
		}
	} else if r.StatusCode() == 500 {
		return nil, xerr.WithCode(xerr.ErrHttpClient, fmt.Sprintf("yxy response: %v", r))
	}

	return &types.GetBusQrcodeResp{
		Qrcode: yxyResp.Qrcode,
	}, nil
}
