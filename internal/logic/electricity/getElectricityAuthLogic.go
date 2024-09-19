package electricity

import (
	"context"

	"yxy-go/internal/svc"
	"yxy-go/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetElectricityAuthLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetElectricityAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetElectricityAuthLogic {
	return &GetElectricityAuthLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetElectricityAuthLogic) GetElectricityAuth(req *types.GetElectricityAuthReq) (resp *types.GetElectricityAuthResp, err error) {
	// todo: add your logic here and delete this line

	return
}
