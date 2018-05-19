package service

import (
	"io"
	"micro/game/pb"

	"golang.org/x/net/context"
	log "github.com/sirupsen/logrus"
)

type GameService struct {
	logger *log.Logger
}

func NewGameService(logger *log.Logger) *GameService {
	return &GameService{
		logger: logger,
	}
}

func (g *GameService) Stream(ctx context.Context, gs pb.GameService_StreamStream) error {
	g.logger.Info("game stream start")
	dieChan := make(chan struct{})
	defer func() {
		close(dieChan)
	}()
	recvChan := g.recv(gs, dieChan)
}

func (g *GameService) recv(stream pb.GameService_StreamStream, dieChan chan struct{}) chan *pb.Frame {
	g.logger.Info("game service grpc recv start")
	recvChan := make(chan *pb.Frame, 1)
	go func() {
		defer func() {
			close(recvChan)
		}()
		for {
			frame, err := stream.Recv()
			if err == io.EOF || err != nil {
				g.logger.Errorf("stream recv err, exit recv loop")
				return
			}
			g.logger.Infof("stream recved:[frame:%s]", frame.Payload)
			select {
			case recvChan <- frame:
			case <-dieChan:
			}
		}
	}()
	return recvChan
}
