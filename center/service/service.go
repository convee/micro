package service

import (
	"encoding/json"
	"micro/center/cache"
	"micro/center/model"
	"micro/center/pb"

	"golang.org/x/net/context"
)

//CenterService 用户中心服务
type CenterService struct {
}

//Auth token验证服务
func (c *CenterService) Auth(ctx context.Context, in *pb.Auth_AuthReq, out *pb.Auth_AuthResp) error {
	newCache := cache.NewCache()
	tokenInfo, err := newCache.GetTokenInfo(in.Token)
	if err != nil {
		return ErrorTokenNotExists
	}
	*out = pb.Auth_AuthResp{TokenInfo: tokenInfo}
	return nil
}

//Login 登录服务
func (c *CenterService) Login(ctx context.Context, in *pb.Login_LoginReq, out *pb.Login_LoginResp) error {
	newModel := model.NewModel()
	user, err := newModel.GetUser(0, in.Username)
	if err != nil {
		return ErrorUserNotExists
	}
	if in.Password != user.Password {
		return ErrorPassword
	}
	token := genToken(in.AppId, in.Username)
	tokenInfo, _ := json.Marshal(&user)

	newCache := cache.NewCache()
	res, err := newCache.SetTokenInfo(token, string(tokenInfo))
	if err != nil || !res {
		return ErrorTokenAdd
	}
	*out = pb.Login_LoginResp{Token: token}
	return nil
}

//Logout 登出服务
func (c *CenterService) Logout(ctx context.Context, in *pb.Login_LogoutReq, out *pb.Login_LogoutResp) error {
	newCache := cache.NewCache()
	err := newCache.DelToken(in.Token)
	if err != nil {
		return ErrorTokenDel
	}
	return nil
}

//GetUser 获取用户信息服务
func (c *CenterService) GetUser(ctx context.Context, in *pb.User_UserReq, out *pb.User_UserResp) error {
	newModel := model.NewModel()
	user, err := newModel.GetUser(in.Uid, in.Username)
	if err != nil {
		return err
	}
	*out = pb.User_UserResp{Uid: user.UID, Username: user.Username}
	return nil
}

//AddUser 添加用户信息服务
func (c *CenterService) AddUser(ctx context.Context, in *pb.User_AddUserReq, out *pb.User_AddUserResp) error {
	newModel := model.NewModel()
	user, err := newModel.GetUser(0, in.Username)
	if user != nil {
		return ErrorUserExists
	}
	uid, err := newModel.CreateUser(in.AppId, in.Username, in.Password, in.Salt)
	if err != nil {
		return ErrorUserAdd
	}
	*out = pb.User_AddUserResp{Uid: int32(uid)}
	return nil
}

//UpdateUser 更新用户信息服务
func (c *CenterService) UpdateUser(ctx context.Context, in *pb.User_UpdateUserReq, out *pb.User_UpdateUserResp) error {
	if in.Uid <= 0 {
		return ErrorParams
	}
	newModel := model.NewModel()
	err := newModel.UpdateUser(in.Uid, in.Username, in.Password)
	if err != nil {
		return ErrorUserUpdate
	}
	return nil
}
