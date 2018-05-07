package server

import (
	conf_struct "config/conf_struct"
	"errors"
	"fmt"
	"github.com/zhangdeman/go-framework/core/conf"
	//"github.com/zhangdeman/go-framework/core/log"
	"net/http"
	"os"
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
		Scheme:      "http",
		AllowIpList: []string{},
		AllowMethod: []string{},
		Env:         "",
		FuncMap:     make(map[string]func() map[string]interface{}),
	}
}

//config
type NewServerConfig struct {
	Scheme      string                                   //请求协议
	AllowIpList []string                                 //允许请求的ip列表
	AllowMethod []string                                 //允许的请求方法
	ListenPort  string                                   //监听的端口
	Env         string                                   //运行环境
	FuncMap     map[string]func() map[string]interface{} //监听的方法
	ConfigPath  string									 //配置路径
}

/**
 * 接口
 */
type NewServerInterface interface {
	MakeServer(configPath string, scheme string, allowIpList []string, allowMethod []string, listenPort string) NewServer //创建一个服务器
	RunServer(env string, configPath string)                                                                       //运行服务器
	AddUriMap(uri string, dealFunc func() map[string]interface{})                                      //增加一个请求map
}

/**
 * 实现 NewServerInterface 接口的结构体
 */
type NewServer struct {
}

/**
 * 创建sever
 */
func (newServer NewServer) MakeServer(configPath string, scheme string, allowIpList []string, allowMethod []string, listenPort string) NewServer {
	NewServerConfigInstance.Scheme = scheme
	NewServerConfigInstance.AllowIpList = allowIpList
	NewServerConfigInstance.AllowMethod = allowMethod
	NewServerConfigInstance.ListenPort = ":" + listenPort
	mapFunc := make(map[string]func() map[string]interface{})
	NewServerConfigInstance.FuncMap = mapFunc

	//设置运行环境
	conf.LoadConfigPath(configPath)
	envConfig , err := conf.LoadConfig("base.yaml", &conf_struct.BaseYaml{})
	if nil != err {
		fmt.Println("base..yaml配置初始化失败", err)
		os.Exit(-1)
	}
	env := conf_struct.GetBaseYamlConfig(envConfig)
	NewServerConfigInstance.Env = env.Env
	NewServerConfigInstance.ConfigPath = configPath

	fmt.Println("初始化日志配置")
	//log.MakeLog()
	return NewServerInstance
}

/**
 * 增加新的请求map
 */
func (newServer NewServer) AddUriMap(uri string, dealFunc func() map[string]interface{}) {
	NewServerConfigInstance.FuncMap[uri] = dealFunc
}

/**
 * 增加新的请求map
 */
func (newServer NewServer) GetUriMap(uri string) func() map[string]interface{} {
	return NewServerConfigInstance.FuncMap[uri]
}

/**
 * 运行server
 */
func (newServer NewServer) RunServer(env string, configPath string) {
	NewServerConfigInstance.Env = env
	fmt.Println("服务器监听端口 " + NewServerConfigInstance.ListenPort)
	for uri, _ := range NewServerConfigInstance.FuncMap {
		fmt.Println("注册请求 : " + uri)
		http.HandleFunc(uri, ResponseData)
	}
	err := http.ListenAndServe(NewServerConfigInstance.ListenPort, nil)
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
func MakeServer(configPath string, scheme string, allowIpList []string, allowMethod []string, listenPort string) NewServer {
	return NewServerInstance.MakeServer(configPath, scheme, allowIpList, allowMethod, listenPort)
}

/**
 * 运行服务器
 */
func RunServer(env string, configPath string) {
	NewServerInstance.RunServer(env, configPath)
}

/**
 * 增加请求map
 */
func AddUriMap(uri string, dealFunc func() map[string]interface{}) {
	NewServerInstance.AddUriMap(uri, dealFunc)
}

func GetUriMap(uri string) (func() map[string]interface{}, error) {
	if _, ok := NewServerConfigInstance.FuncMap[uri]; ok {
		//存在
		return NewServerConfigInstance.FuncMap[uri], nil
	}
	return nil, errors.New("请求未注册")
}
