package boot

import (
	"github.com/zhangdeman/go-framework/core/server"
	"fmt"
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
	//newServer.AddUriMap()
	RunConfigInstance.RunServer = newServer
	//运行服务器
	RunConfigInstance.RunServer.RunServer()
}

/**
 * 运行引导接口
 */
type RunBootInterface interface {

}

/**
 * 运行引导结构体
 */
type RunBoot struct {

}

/**
 * 运行之后的一些配置
 */
type RunConfig struct {
	RunServer server.NewServerInterface
}