package electricity

import (
	"context"

	"yxy-go/internal/svc"
	"yxy-go/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetElectricityUsageRecordsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetElectricityUsageRecordsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetElectricityUsageRecordsLogic {
	return &GetElectricityUsageRecordsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetElectricityUsageRecordsLogic) GetElectricityUsageRecords(req *types.GetElectricityUsageRecordsReq) (resp *types.GetElectricityUsageRecordsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
