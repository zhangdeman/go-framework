package boot

import (
	"fmt"
	"github.com/zhangdeman/go-framework/core/server"
	"router"
)

var (
	RunConfigInstance RunConfig
)

/**
 * 初始化函数
 */
func init() {

}

/**
 * 运行服务器
 */
func RunServer(env string, configPath string) {
	RunConfigInstance = RunConfig{}
	//初始化服务器
	newServer := server.MakeServer(configPath, "http", []string{}, []string{}, "8990")
	fmt.Println(newServer)
	//初始化配置
	for uri, method := range router.RouterMap {
		newServer.AddUriMap(uri, method)
	}
	RunConfigInstance.RunServer = newServer
	//运行服务器
	RunConfigInstance.RunServer.RunServer(env, configPath)
}

/**
 * 运行之后的一些配置
 */
type RunConfig struct {
	RunServer server.NewServerInterface
}
