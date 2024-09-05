package login

import (
	"context"

	"yxy-go/internal/svc"
	"yxy-go/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSecurityTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSecurityTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSecurityTokenLogic {
	return &GetSecurityTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSecurityTokenLogic) GetSecurityToken(req *types.GetSecurityTokenReq) (resp *types.GetSecurityTokenResp, err error) {
	// todo: add your logic here and delete this line

	return
}
