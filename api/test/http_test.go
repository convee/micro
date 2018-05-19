package test

import (
	"net/http"
	"net/url"
	"log"
	"io/ioutil"
	"fmt"
	"testing"
)

func Test_Login(t *testing.T) {
	data := url.Values{"username":{"wangkang"}, "password":{"wang1019"}, "app_id":{"1"}}
	resp, err := http.PostForm("http://127.0.0.1:8001/login", data)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))

}

func Test_Register(t *testing.T) {
	data := url.Values{"username":{"wangkang"}, "password":{"wang1019"}, "app_id":{"1"}}
	resp, err := http.PostForm("http://127.0.0.1:8001/register", data)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))

}