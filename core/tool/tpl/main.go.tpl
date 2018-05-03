package main

import (
	"github.com/zhangdeman/go-framework/core/boot"
	"github.com/zhangdeman/go-framework/core/conf"
	"config/conf_struct"
	"fmt"
	"reflect"
)

func init()  {
	//注册配置文件路径
	conf.LoadConfigPath("./src/config")
	data, _ := conf.LoadConfig("base.yaml", &conf_struct.BaseYaml{})

	value := reflect.ValueOf(data)
	baseConfig := conf_struct.GetBaseYamlConfig(data)

	fmt.Println("fghjhkgjjkkjkghjghj",baseConfig.Host)

	fmt.Println("main",value,reflect.TypeOf(data))

	//启动运行服务器
	boot.RunServer()
}

func main()  {

}
