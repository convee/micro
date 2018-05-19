package test

import (
	"micro/game/pb"
	"testing"

	"github.com/micro/go-grpc"
	"golang.org/x/net/context"
)

func Test_Stream(t *testing.T) {
	grpcService := grpc.NewService()
	grpcService.Init()
	client := pb.NewGameServiceClient("go.micro.srv.game", grpcService.Client())
	stream, err := client.Stream(context.Background())
	if err != nil {
		t.Fatalf("failed to call: %v", err)
	}
	for {
		err := stream.Send(&pb.Frame{Payload: []byte("6666666")})
		if err != nil {
			t.Fatalf("failed to send: %v", err)
			break
		}
		reply, err := stream.Recv()
		if err != nil {
			t.Fatalf("failed to recv: %v", err)
			break
		}
		t.Logf("reply:, %v", reply)
	}
}
