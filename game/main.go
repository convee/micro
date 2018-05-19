package main

import (
	"micro/game/pb"
	"micro/game/service"

	"github.com/micro/go-grpc"
	"github.com/micro/go-micro"
	log "github.com/sirupsen/logrus"
)

func main() {
	//日志初始化
	var logger = log.New()
	logger.WithFields(log.Fields{
		"log_name": "game",
	})

	gSrv := grpc.NewService(
		micro.Name("go.micro.srv.game"),
	)
	gSrv.Init()
	logger.Info("start a game service")
	gameService := service.NewGameService(logger)
	pb.RegisterGameServiceHandler(gSrv.Server(), gameService)
	if err := gSrv.Run(); err != nil {
		logger.Errorf("error: %v", err)
	}
}
