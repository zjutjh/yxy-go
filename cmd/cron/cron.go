package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"yxy-go/internal/config"
	"yxy-go/internal/cron"
	"yxy-go/internal/svc"
)

var configFile = flag.String("f", "etc/yxy-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	cronJob := cron.NewCronJob(context.Background(), ctx)
	cronJob.Register()

	fmt.Printf("Starting cron service...\n")
	ctx.Cron.Start()
	defer ctx.Cron.Stop()

	logx.DisableStat()

	select {} // Keep the service running
}
