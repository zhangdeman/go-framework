package conf

import (
	"io/ioutil"
	"fmt"
	"gopkg.in/yaml.v2"
)
type ConfigConf struct {
	configPath    string                 //配置文件存储路径
	hasLoadConfigMap map[string][]byte	//已加载配置文件
}

type ConfigInterface interface {
	LoadConfigPath(configPath string) //加载配置文件
	LoadConfig(configName string, data interface{}) (interface{}, error)     //配置文件路径下必须有base.yaml配置文件,定义基础公用的配置
	HasLoadConfig(fileName string) ([]byte, bool)	//获取已加载配置文件
}

type Config struct {
}

var (
	ConfigConfInstance ConfigConf
	ConfigInstance     ConfigInterface
)

/**
 * 初始化函数
 */
func init() {
	ConfigConfInstance = ConfigConf{
		configPath:"",
		hasLoadConfigMap:make(map[string][]byte),
	}
	ConfigInstance = Config{}
}

/**
 * 加载配置文件路径
 */
func (c Config) LoadConfigPath(configPath string) {
	ConfigConfInstance.configPath = configPath
}

/**
 * 加载公共配置文件
 */
func (c Config) LoadConfig(configFileName string, data interface{}) (interface{},error) {
	var yamlFile []byte
	var err error
	var ok bool
	yamlFile, ok = c.HasLoadConfig(configFileName)
	if !ok {
		//配置没加载过，进行加载，已加载过，不作任何处理
		yamlFile, err = ioutil.ReadFile(ConfigConfInstance.configPath+"/"+configFileName)
		fmt.Println(string(yamlFile))
		if err != nil {
			fmt.Println("load config error", err)
			return nil, err
		}
		ConfigConfInstance.hasLoadConfigMap[configFileName] = yamlFile
		err := yaml.Unmarshal(yamlFile,data)
		if err != nil {
			return nil, err
		}
	}
	return data, nil
}

/**
 * 是否已加载配置文件
 */
func (c Config) HasLoadConfig(fileName string) ([]byte, bool)  {
	yamlFile, ok := ConfigConfInstance.hasLoadConfigMap[fileName]
	if ok {
		return yamlFile, ok
	}
	return nil, false
}

/**
 * 加载配置文件路径
 */
func LoadConfigPath(configPath string) {
	ConfigInstance.LoadConfigPath(configPath)
}

/**
 * 加载公共配置文件
 */
func LoadConfig(configFileName string, data interface{}) (interface{}, error) {
	return ConfigInstance.LoadConfig(configFileName, data)
}

/**
 * 是否已加载配置文件
 */
func HasLoadConfig(fileName string) ([]byte, bool) {
	return ConfigInstance.HasLoadConfig(fileName)
}
