package server

import (
	"fmt"
	"net/http"
	"errors"
	"github.com/zhangdeman/go-framework/core/conf"
	"os"
	conf_struct "config/conf_struct"
)

/**
 * 创建服务器，监听http请求
 */

var (
	NewServerInstance       NewServer
	NewServerConfigInstance NewServerConfig
)

func init() {
	NewServerInstance = NewServer{}
	NewServerConfigInstance = NewServerConfig{
		scheme:"http",
		allowIpList:[]string{},
		allowMethod:[]string{},
		env:"",
		funcMap:make(map[string] func()map[string]interface{}),
	}
}

//config
type NewServerConfig struct {
	scheme      string            //请求协议
	allowIpList []string          //允许请求的ip列表
	allowMethod []string          //允许的请求方法
	listenPort  string            //监听的端口
	env 		string			  //运行环境
	funcMap     map[string]func() map[string]interface{} //监听的方法
}

/**
 * 接口
 */
type NewServerInterface interface {
	MakeServer(scheme string, allowIpList []string, allowMethod []string, listenPort string) NewServer	//创建一个服务器
	RunServer(configPath string)	//运行服务器
	AddUriMap(uri string, dealFunc func() map[string]interface{})	//增加一个请求map
}

/**
 * 实现 NewServerInterface 接口的结构体
 */
type NewServer struct {
}

/**
 * 创建sever
 */
func (newServer NewServer) MakeServer(scheme string, allowIpList []string, allowMethod []string, listenPort string) NewServer {
	NewServerConfigInstance.scheme = scheme
	NewServerConfigInstance.allowIpList = allowIpList
	NewServerConfigInstance.allowMethod = allowMethod
	NewServerConfigInstance.listenPort = ":"+listenPort
	mapFunc := make(map[string]func()map[string]interface{})
	NewServerConfigInstance.funcMap = mapFunc
	return NewServerInstance
}

/**
 * 增加新的请求map
 */
func (newServer NewServer) AddUriMap(uri string, dealFunc func() map[string]interface{})  {
	NewServerConfigInstance.funcMap[uri] = dealFunc
}

/**
 * 增加新的请求map
 */
func (newServer NewServer) GetUriMap(uri string) func() map[string]interface{} {
	return NewServerConfigInstance.funcMap[uri]
}

/**
 * 运行server
 */
func (newServer NewServer) RunServer(configPath string) {
	//加载配置文件
	conf.LoadConfigPath(configPath)
	config, err := conf.LoadConfig("base.yaml", &conf_struct.BaseYaml{})
	if nil != err {
		fmt.Println("配置文件 base.yaml 加载失败",err)
		os.Exit(-1)
	}
	envConfig := conf_struct.GetBaseYamlConfig(config)
	NewServerConfigInstance.env = envConfig.Env
	fmt.Println("设置运行环境 : "+NewServerConfigInstance.env)


	fmt.Println("服务器监听端口 " + NewServerConfigInstance.listenPort)
	for uri, _ := range NewServerConfigInstance.funcMap{
		fmt.Println("注册请求 : " + uri )
		http.HandleFunc(uri, ResponseData)
	}
	err = http.ListenAndServe(NewServerConfigInstance.listenPort, nil)
	if err != nil {
		fmt.Println(err)
	}

}

func (newServer NewServer) ValidateRequestIp() {

}

func (newServer NewServer) ValidateRequestMethod() {

}

/**
 * 创建服务器
 */
func MakeServer(scheme string, allowIpList []string, allowMethod []string, listenPort string) NewServer {
	return NewServerInstance.MakeServer(scheme, allowIpList, allowMethod, listenPort)
}

/**
 * 运行服务器
 */
func RunServer(configPath string)  {
	NewServerInstance.RunServer(configPath)
}

/**
 * 增加请求map
 */
func AddUriMap(uri string, dealFunc func() map[string]interface{}) {
	NewServerInstance.AddUriMap(uri, dealFunc)
}

func GetUriMap(uri string) (func()map[string]interface{}, error) {
	if _, ok := NewServerConfigInstance.funcMap[uri]; ok {
		//存在
		return NewServerConfigInstance.funcMap[uri], nil
	}
	return nil,errors.New("请求未注册")
}
