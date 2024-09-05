package card

import (
	"context"

	"yxy-go/internal/svc"
	"yxy-go/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCardConsumptionRecordsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCardConsumptionRecordsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCardConsumptionRecordsLogic {
	return &GetCardConsumptionRecordsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCardConsumptionRecordsLogic) GetCardConsumptionRecords(req *types.GetCardConsumptionRecordsReq) (resp *types.GetCardConsumptionRecordsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
