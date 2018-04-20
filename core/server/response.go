package server

import (
	"net/http"
	"fmt"
)

var (
	ResponseInstance Response
	DoResponseInstance DoResponseInterface
)

type DoResponseInterface interface{
	ResponseData(w http.ResponseWriter, r *http.Request)
}

type Response struct {
	w http.ResponseWriter
	r *http.Request
}


/**
 * 响应结构体
 */
type DoResponse struct {

}

func init()  {
	ResponseInstance = Response{}
	DoResponseInstance = DoResponse{}
	//ResponseData()
}

func (resp DoResponse) ResponseData(w http.ResponseWriter, r *http.Request) {
	uri := r.RequestURI
	dealFunc := GetUriMap(uri)
	data := dealFunc()
	fmt.Fprint(w, data)
}

func (resp DoResponse) GetDealFunc(uri string, server NewServerInterface) {
	fmt.Fprint(ResponseInstance.w, []string{})
}

func ResponseData(w http.ResponseWriter, r *http.Request)  {
	DoResponseInstance.ResponseData(w, r)
}
