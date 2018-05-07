package log

import (
	conf_struct "config/conf_struct"
	"github.com/zhangdeman/go-framework/core/conf"
	"fmt"
	"os"
)

type LogInterface interface {
	MakeLog(env string) LogInterface
	Trace(data interface{})
	Debug(data interface{})
	Notice(data interface{})
	Warn(data interface{})
	Fatal(data interface{})
	LogData(data interface{}, logLevel uint)
	IsAllowLog() bool
}

const (
	LogLevelTrace = 1	//trace级别
	LogLevelDebug = 2	//debug级别
	LogLevelNotice= 3	//notice级别
	LogLevelWarn  = 4	//warn级别
	LogLevelFatal = 5   //fatal级别
)

var (
	LogConfigInstance LogConfig
	LogInstance       LogInterface
)

type LogConfig struct {
	LogEnv   string //运行环境
	LogLevel uint   //日志级别
	LogPath  string //日志路径
}

type Log struct {
}

/**
 * 初始化log
 */
func (log Log) MakeLog(env string) LogInterface  {

	fmt.Println("运行环境 : zhang", env)
	/*LogConfigInstance = LogConfig{}
	LogInstance = &Log{}
	LogConfigInstance.LogEnv = env
	return LogInstance*/
	return  nil
}

/**
 * 记录trace日志
 */
func (log Log) Trace(logData interface{}) {
	log.LogData(logData, LogLevelTrace)
}

/**
 * 记录notice日志
 */
func (log Log) Notice(logData interface{}) {
	log.LogData(logData, LogLevelNotice)
}

/**
 * 记录trace日志
 */
func (log Log) Debug(logData interface{}) {
	log.LogData(logData, LogLevelDebug)
}

/**
 * 记录trace日志
 */
func (log Log) Warn(logData interface{}) {
	log.LogData(logData, LogLevelWarn)
}

/**
 * 记录trace日志
 */
func (log Log) Fatal(logData interface{}) {
	log.LogData(logData, LogLevelFatal)
}

/**
 * 统一记录日志的方法
 */
func (log Log) LogData(logData interface{}, logLevel uint)  {
	fmt.Println("logData", logData)
	if LogConfigInstance.LogLevel <= logLevel {
		//加载配置文件
		config, err := conf.LoadConfig("base.yaml", conf_struct.BaseYaml{})
		if nil != err {
			fmt.Println("配置文件 base.yaml 加载失败")
			os.Exit(-1)
		}
		fmt.Println(config, "log init")
	}
}

/**
 * 判断是否允许记录当前级别日志
 */
func (log Log) IsAllowLog() (bool)  {
	return true
}

func MakeLog(env string) LogInterface  {
	fmt.Println("debug env", env)
	return LogInstance.MakeLog(env)
}

func Trace(data interface{})  {
	LogInstance.Trace(data)
}
