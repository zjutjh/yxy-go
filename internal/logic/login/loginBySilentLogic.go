package login

import (
	"context"

	"yxy-go/internal/svc"
	"yxy-go/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginBySilentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginBySilentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginBySilentLogic {
	return &LoginBySilentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginBySilentLogic) LoginBySilent(req *types.LoginBySilentReq) (resp *types.LoginBySilentResp, err error) {
	// todo: add your logic here and delete this line

	return
}
