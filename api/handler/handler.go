package handler

import (
	"net/http"
	"encoding/json"
	"micro/api/client"
	"micro/api/client/pb"
	"strings"
	"strconv"
)

type Handler struct {}

type Data  map[string]interface{}

type Response struct {
	ErrorCode int `json:"error_code"`
	ErrorMsg string `json:"error_msg"`
	Data Data `json:"data"`
}

type LoginReq struct {
	AppId int32 `json:"app_id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResp struct {
	Token string `json:"token"`
}

func (h Handler) Login(w http.ResponseWriter, r *http.Request)  {

	r.ParseForm()
	username := strings.TrimSpace(r.Form.Get("username"))
	password := strings.TrimSpace(r.Form.Get("password"))
	app_id,_ := strconv.Atoi(r.Form.Get("app_id"))
	if username == "" || password == "" || app_id <= 0 {
		json.NewEncoder(w).Encode(Response{ErrorCode:1, ErrorMsg:"参数错误"})
		return
	}
	cli := client.NewClient()

	req := &pb.Login_LoginReq{AppId:int32(app_id), Username:username, Password:password}
	resp, err := cli.Login(req)
	if err != nil {
		json.NewEncoder(w).Encode(Response{ErrorCode:1, ErrorMsg:"登录失败，请重试"})
		return
	}
	token := resp.Token
	data := Data{"token":token}
	json.NewEncoder(w).Encode(Response{ErrorCode:1, Data:data})
	return

}

func (h Handler) Logout(w http.ResponseWriter, r *http.Request)  {

}

