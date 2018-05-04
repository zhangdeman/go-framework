package log

import (
	"github.com/zhangdeman/go-framework/core/server"
	"fmt"
)

type LogInterface interface {
	
}

const (
	LogLevelDebug = 1
)

var (
	LogConfigInstance LogConfig
	LogInstance LogInterface
)

type LogConfig struct {
	LogEnv	string	//运行环境
	LogLevel uint	//日志级别
	LogPath  string	//日志路径
}

type Log struct {

}

func init()  {
	//加载配置文件
	/*config, err := conf.LoadConfig("base.yaml", conf_struct.BaseYaml{})
	if nil != err {
		fmt.Println("配置文件 base.yaml 加载失败")
		os.Exit(-1)
	}*/
	fmt.Println(server.NewServerConfigInstance.env)
}

func Test()  {
	
}
