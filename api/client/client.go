package client

import (
	"micro/api/client/pb"
	"github.com/micro/go-grpc"
	"golang.org/x/net/context"
)

type Client struct {
	gClient pb.CenterServiceClient
}

func NewClient() *Client {
	gSrv := grpc.NewService()
	gSrv.Init()
	return &Client{
		gClient: pb.NewCenterServiceClient("go.micro.srv.center", gSrv.Client()),
	}
}

//获取用户信息
func (c *Client)GetUser(req *pb.User_UserReq) (*pb.User_UserResp, error) {
	userInfo, err := c.gClient.GetUser(context.Background(), req)
	if err != nil {
		return nil, err
	}
	return userInfo, nil
}


func (c *Client)Login(req *pb.Login_LoginReq) (*pb.Login_LoginResp, error) {
	loginResp, err := c.gClient.Login(context.Background(), req)
	if err != nil {
		return nil, err
	}
	return loginResp, nil
}