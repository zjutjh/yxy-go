package electricity

import (
	"context"

	"yxy-go/internal/svc"
	"yxy-go/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetElectricitySurplusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetElectricitySurplusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetElectricitySurplusLogic {
	return &GetElectricitySurplusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetElectricitySurplusLogic) GetElectricitySurplus(req *types.GetElectricitySurplusReq) (resp *types.GetElectricitySurplusResp, err error) {
	// todo: add your logic here and delete this line

	return
}
