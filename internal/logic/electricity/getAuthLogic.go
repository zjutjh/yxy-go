package electricity

import (
	"context"

	"yxy-go/internal/svc"
	"yxy-go/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAuthLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAuthLogic {
	return &GetAuthLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAuthLogic) GetAuth(req *types.GetAuthReq) (resp *types.GetAuthResp, err error) {
	// todo: add your logic here and delete this line

	return
}
