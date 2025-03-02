package cron

import (
	"context"
	"yxy-go/internal/svc"
)

type CronJob struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCronJob(ctx context.Context, svcCtx *svc.ServiceContext) *CronJob {
	return &CronJob{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (c *CronJob) Register() {
	_, err := c.svcCtx.Cron.AddFunc(c.svcCtx.Config.CronTime, func() {
		l := NewSendLowBatteryAlertLogic(c.ctx, c.svcCtx)
		l.Logger.Info("Start sending low battery alert")
		l.SendLowBatteryAlertLogic()
		l.Logger.Info("Finish sending low battery alert")
	})
	if err != nil {
		panic(err)
	}
}
