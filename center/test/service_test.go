package test

import (
	"crypto/md5"
	"fmt"
	"micro/center/pb"
	"testing"

	grpc "github.com/micro/go-grpc"
	"golang.org/x/net/context"
)

func NewClient() pb.CenterServiceClient {
	gSrv := grpc.NewService()
	gSrv.Init()
	return pb.NewCenterServiceClient("go.micro.srv.center", gSrv.Client())
}

func Md5(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

//添加用户
func Test_AddUser(t *testing.T) {
	salt := "123456"
	password := Md5(Md5("wang1019") + salt)
	client := NewClient()
	in := &pb.User_AddUserReq{
		AppId:    1,
		Username: "convee",
		Password: password,
		Salt:     salt,
	}
	t.Logf("%v", in)
	rs, err := client.AddUser(context.Background(), in)
	if err != nil {
		t.Fatalf("AddUser Error: %v", err)
	}
	t.Logf("AddUser:[%v]", rs)

}
