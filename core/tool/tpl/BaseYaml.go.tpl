package conf_struct

type BaseYaml struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	User string `yaml:"user"`
	Pwd string `yaml:"pwd"`
	Dbname string `yaml:"dbname"`
	Env string `yaml:"env"`
}

func GetBaseYamlConfig(data interface{}) *BaseYaml  {
	var baseConfig  *BaseYaml
	switch config := data.(type) {
	case *BaseYaml:
		baseConfig = config
	default:
		return nil
	}
	return baseConfig
}
