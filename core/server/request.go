package server

type ClientRequestInterface interface {
}

type ClientRequestInfo struct {
	method string //请求方法
}

var (
	ClientRequestInfoInstance ClientRequestInfo
	ClientRequestInstance     ClientRequest
)

func init() {
	ClientRequestInfoInstance = ClientRequestInfo{}
	ClientRequestInstance = ClientRequest{}
}

type ClientRequest struct {
}

func (client ClientRequest) RequestMethod() {

}
