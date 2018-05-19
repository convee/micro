package main

import (
	"micro/center/pb"
	"micro/center/service"
	"time"

	log "github.com/sirupsen/logrus"
	grpc "github.com/micro/go-grpc"
	"github.com/micro/go-micro"
	"gitlab.com/convee/system"
)

func main() {
	logger := log.WithFields(log.Fields{
		"srv": "center",
	})
	logger.Info("go.micro.srv.center start")
	gSrv := grpc.NewService(
		micro.Name("go.micro.srv.center"),
		micro.Version("1.0"),
		micro.RegisterTTL(time.Second),
	)
	gSrv.Init()
	system.Run("config.toml")

	pb.RegisterCenterServiceHandler(gSrv.Server(), new(service.CenterService))
	if err := gSrv.Run(); err != nil {
		logger.Errorf("gSrv Run Error: %v", err)
	}

}
