package main

import (
	"fmt"
	"bufio"
	"os"
	"os/exec"
	"bytes"
	"io/ioutil"
	"encoding/json"
	"strings"
)

var supportCmd []string

func init()  {
	supportCmd = []string{
		"create_project",
		"create_config",
		"delete_config",
		"run",
	}
}

//创建项目的配置
var createProjectConfig map[string]string

/**
 * 创建一个 go 项目
 */
func CreateProject()  {
	createProjectConfig = make(map[string]string)

	remindMsg := make(map[string]string)
	remindMsg["store_path"] = "请输入项目存储路径 : "
	remindMsg["project_name"] = "请输入项目名 : "
	remindMsg["project_desc"] = "请输入项目描述 : "

	remindKeyList := []string{
		"store_path",
		"project_name",
		"project_desc",
	}

	remindMsglen := len(remindKeyList)

	var key string
	for i := 0; i < remindMsglen; i++ {
		key = remindKeyList[i]
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(remindMsg[key])
		data, _, _ := reader.ReadLine()
		createProjectConfig[key] = string(data)
	}

	storePath := createProjectConfig["store_path"]	//存储路径
	projectPath := storePath+"/"+createProjectConfig["project_name"]	//项目路径

	//创建目录
	var err error
	err = os.MkdirAll(storePath, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	fmt.Println("创建项目存储目录 : "+ storePath + " 成功! ")
	//创建项目
	_, err = ExecShell("cd "+storePath)
	if err != nil {
		fmt.Println("无法切换至项目存储目录")
		os.Exit(-2)
	}
	fmt.Println("切换至项目存储目录成功, 当前位置 : " + storePath)
	err = os.MkdirAll(projectPath, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		os.Exit(-3)
	}
	fmt.Println("项目创建成功")
	if err != nil {
		fmt.Println("无法切换至项目目录")
		os.Exit(-4)
	}
	fmt.Println("切换至项目目录成功, 当前位置 : " + projectPath)

	/**
	 * |-- 项目目录
	 *    |-- src
	          |-- config 配置文件目录
				 |-- config_struct 存储配置文件对应的结构体
					|-- baseYaml.go
				 |--base.yaml
	          |-- route 路由配置
	          |-- controller  业务入口控制器
	          |-- model   模型层
		mian.go 入口文件
	 */
	//初始化项目结构
	//切换至项目目录
	_, err = ExecShell("cd "+projectPath)

	//创建描述文件
	descFile, err := os.OpenFile("README.md", os.O_RDWR|os.O_CREATE, 0766)
	if err != nil {
		fmt.Println("描述文件创建失败!")
		os.Exit(-4)
	}
	fmt.Println("描述文件创建成功!")
	descFile.Close()

	//初始化项目
	//定义需要创建的目录
	needCreateDirList := []string{
		"src/config/conf_struct",
		"src/route",
		"src/controller",
		"src/model",
		"src/router",
	}

	dirLen := len(needCreateDirList)
	for index := 0; index < dirLen; index++ {
		err = os.MkdirAll(projectPath+"/"+needCreateDirList[index], os.ModePerm)
		if err != nil {
			fmt.Println("初始化项目失败 !")
			os.Exit(-5)
		}
	}

	//导入GOPATH
	_,err = ExecShell("export GOPATH="+projectPath)
	if err != nil {
		fmt.Println("导入GOPATH失败")
		os.Exit(-6)
	}

	//解决依赖
	dependence, err := os.Open("dependence.json")
	if err != nil {
		fmt.Println("处理依赖失败，请手动处理缺少的library")
	} else {
		dep, _ := ioutil.ReadAll(dependence)
		var depMap map[string]string
		json.Unmarshal(dep, &depMap)
		for key, addr := range depMap{
			fmt.Print("解决依赖 "+key +" : ")
			ExecShell("go get "+addr)
			fmt.Println("done")
		}
	}

	//main.go
	err = WriteFile("tpl/main.go.tpl", projectPath+"/main.go")
	if err != nil {
		os.Exit(-7)
	}

	//base.yaml
	err = WriteFile("tpl/base.yaml.tpl", projectPath+"/src/config/base.yaml")
	if err != nil {
		os.Exit(-7)
	}

	//BaseYaml.go
	err = WriteFile("tpl/BaseYaml.go.tpl", projectPath+"/src/config/conf_struct/BaseYaml.go")
	if err != nil {
		os.Exit(-7)
	}

	//TestController
	err = WriteFile("tpl/TestController.go.tpl", projectPath+"/src/controller/TestController.go")
	if err != nil {
		os.Exit(-7)
	}

	//route.go
	err = WriteFile("tpl/route.go.tpl", projectPath+"/src/route/route.go")
	if err != nil {
		os.Exit(-7)
	}

	fmt.Println("初始化项目成功")

}

func CreateConfig(projectPath string)  {

}

/**
 * 运行一个项目
 */
func Run(projectName string)  {

}

func main()  {
	cmdParams := os.Args
	cmdParamsLen := len(cmdParams)
	if(cmdParamsLen < 2) {
		fmt.Println("请选择执行的操作,支持的操作 : "+strings.Join(supportCmd, ","))
		os.Exit(-11)
	}
	cmd := cmdParams[1]
	switch cmd {
	case "create_project":
		CreateProject()
		break
	default:
		fmt.Println("不支持命令 : "+cmd)
	}
}

//阻塞式的执行外部shell命令的函数,等待执行完毕并返回标准输出
func ExecShell(s string) (string, error){
	//函数返回一个*Cmd，用于使用给出的参数执行name指定的程序
	cmd := exec.Command("/bin/bash", "-c", s)

	//读取io.Writer类型的cmd.Stdout，再通过bytes.Buffer(缓冲byte类型的缓冲器)将byte类型转化为string类型(out.String():这是bytes类型提供的接口)
	var out bytes.Buffer
	cmd.Stdout = &out

	//Run执行c包含的命令，并阻塞直到完成。  这里stdout被取出，cmd.Wait()无法正确获取stdin,stdout,stderr，则阻塞在那了
	err := cmd.Run()

	return out.String(), err
}

/**
 * 从模板中读取内容，写入目标文件
 */
func WriteFile(tplFile string, targetFile string) error {
	//入口文件建立
	contentFile, err := os.Open(tplFile)
	if err != nil {
		fmt.Println(tplFile + "文件模板不存在")
		return err
	}

	content , _ := ioutil.ReadAll(contentFile)

	if err != nil {
		fmt.Println(targetFile + "入口文件创建失败!")
		return err
	}
	err = ioutil.WriteFile(targetFile,content, os.ModePerm)
	if err != nil {
		fmt.Println(targetFile + "入口文件创建失败")
		return err
	}
	fmt.Println(targetFile + "文件创建成功")
	return nil
}
