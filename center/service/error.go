package service

import (
	"github.com/micro/go-micro/errors"
)



var (
	//ErrorSystem 系统函数错误
	ErrorSystem = errors.New("go.micro.srv.center", "system error", 1000)

	//ErrorParams 参数错误
	ErrorParams = errors.New("go.micro.srv.center", "params error", 1001)

	//ErrorUserNotExists 用户不存在
	ErrorUserNotExists = errors.New("go.micro.srv.center", "user not exists", 2001)
	//ErrorPassword 密码错误
	ErrorPassword = errors.New("go.micro.srv.center", "password error", 2002)
	//ErrorUserExists 用户已存在
	ErrorUserExists = errors.New("go.micro.srv.center", "user exists", 2003)
	//ErrorUserAdd 用户添加失败
	ErrorUserAdd = errors.New("go.micro.srv.center", "user add error", 2004)
	//ErrorTokenNotExists token不存在
	ErrorTokenNotExists = errors.New("go.micro.srv.center", "token not exists", 2005)
	//ErrorTokenAdd token创建失败
	ErrorTokenAdd = errors.New("go.micro.srv.center", "token add error", 2006)
	//ErrorTokenDel token清除失败
	ErrorTokenDel = errors.New("go.micro.srv.center", "token del error", 2007)
	//ErrorUserUpdate 用户更新失败
	ErrorUserUpdate = errors.New("go.micro.srv.center", "user update error", 2008)
)
