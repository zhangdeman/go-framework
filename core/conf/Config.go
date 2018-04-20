package conf

import (
	"github.com/c4pt0r/ini"
)

type ConfigConf struct {
	configPath string	//配置文件存储路径
	hasLoadConfig map[string]ini.ConfSet	//已经加载的配置,避免重复加载
}


type ConfigInterface interface {
	LoadConfigPath(configPath string)	//加载配置文件
	LoadConfig(configName string)					//配置文件路径下必须有common.ini配置文件,定义基础公用的配置
}

type Config struct {

}

var (
	ConfigConfInstance ConfigConf
	ConfigInstance ConfigInterface
)

/**
 * 初始化函数
 */
func init()  {
	ConfigConfInstance = ConfigConf{}
	ConfigInstance = Config{}
}

/**
 * 加载配置文件路径
 */
func (c Config) LoadConfigPath(configPath string)  {
	ConfigConfInstance.configPath = configPath
}

/**
 * 加载公共配置文件
 */
func (c Config) LoadConfig(configFileName string)  {
	_ , ok := ConfigConfInstance.hasLoadConfig[configFileName]
	if !ok {
		//配置没加载过，进行加载，已加载过，不作任何处理
		conf := ini.NewConf(ConfigConfInstance.configPath+"/"+configFileName)
		ConfigConfInstance.hasLoadConfig[configFileName] = conf
	}
}

/**
 * 加载配置文件路径
 */
func LoadConfigPath(configPath string)  {
	ConfigInstance.LoadConfig(configPath)
}

/**
 * 加载公共配置文件
 */
func LoadConfig(configFileName string)  {
	ConfigInstance.LoadConfig(configFileName)
}


