package electricity

import (
	"context"

	"yxy-go/internal/svc"
	"yxy-go/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetElectricityRechargeRecordsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetElectricityRechargeRecordsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetElectricityRechargeRecordsLogic {
	return &GetElectricityRechargeRecordsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetElectricityRechargeRecordsLogic) GetElectricityRechargeRecords(req *types.GetElectricityRechargeRecordsReq) (resp *types.GetElectricityRechargeRecordsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
