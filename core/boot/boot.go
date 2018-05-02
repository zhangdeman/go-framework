package boot

import (
	"github.com/zhangdeman/go-framework/core/server"
	"fmt"
	"router"
)

var (
	RunConfigInstance RunConfig
)
/**
 * 初始化函数
 */
func init()  {
	RunConfigInstance = RunConfig{}
	//初始化服务器
	newServer := server.MakeServer("http", []string{}, []string{}, "8990")
	fmt.Println(newServer)
	//初始化配置
	for uri, method := range router.RouterMap{
		newServer.AddUriMap(uri,method)
	}
	RunConfigInstance.RunServer = newServer
}

/**
 * 运行服务器
 */
func RunServer()  {
	//运行服务器
	RunConfigInstance.RunServer.RunServer()
}

/**
 * 运行之后的一些配置
 */
type RunConfig struct {
	RunServer server.NewServerInterface
}