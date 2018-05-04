package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	ResponseInstance   Response
	DoResponseInstance DoResponseInterface
	ShowData           interface{}
)

type DoResponseInterface interface {
	ResponseData(w http.ResponseWriter, r *http.Request)
}

type Response struct {
	w    http.ResponseWriter
	r    *http.Request
	data interface{}
}

/**
 * 响应结构体
 */
type DoResponse struct {
}

func init() {
	ResponseInstance = Response{}
	DoResponseInstance = DoResponse{}
	//ResponseData()
}

func (resp DoResponse) ResponseData(w http.ResponseWriter, r *http.Request) {
	uri := r.RequestURI
	fmt.Println("请求uri : " + uri)
	dealFunc, err := GetUriMap(uri)
	if nil != err {
		fmt.Println("未注册请求 : 404")
		//未获取到请求map
		w.WriteHeader(404)
	} else {
		ShowData = dealFunc()
		r.Body.Close()
		str, _ := json.Marshal(ShowData)
		fmt.Println("响应数据 : ", string(str))
		fmt.Fprint(w, string(str))
	}

}

func (resp DoResponse) GetDealFunc(uri string, server NewServerInterface) {
	fmt.Fprint(ResponseInstance.w, []string{})
}

func ResponseData(w http.ResponseWriter, r *http.Request) {
	DoResponseInstance.ResponseData(w, r)
}
