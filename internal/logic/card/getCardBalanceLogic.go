package card

import (
	"context"

	"yxy-go/internal/svc"
	"yxy-go/internal/types"

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

func (l *GetCardBalanceLogic) GetCardBalance(req *types.GetCardBalanceReq) (resp *types.GetCardBalanceResp, err error) {
	// todo: add your logic here and delete this line

	return
}
