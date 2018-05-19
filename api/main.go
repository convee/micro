package main

import (
	"micro/api/handler"
	"net/http"
)

func main()  {
	handle := &handler.Handler{}
	http.HandleFunc("/login", handle.Login)
	http.ListenAndServe(":10000", nil)
}