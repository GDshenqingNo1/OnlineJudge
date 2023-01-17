package boot

import (
	g "OnlineJudge/app/global"
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"os"
)

const (
	configEnv  = "JUDGE_CONFIG_PATH"
	configFile = "manifest/config/config.yaml"
)

func ViperSetup(path ...string) {
	var configPath string
	if len(path) != 0 {
		configPath = path[0] //参数
	} else { //命令行参数
		flag.StringVar(&configPath, "c", "", "set config file path")
		flag.Parse()
		if configPath == "" {
			if configPath = os.Getenv(configEnv); configPath != "" {
				//环境变量
			} else {
				configPath = configFile

			}
		}

	}
	fmt.Printf("get configPath:%s\n", configPath)
	v := viper.New()
	v.SetConfigFile(configPath)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("get config file failed,err:%v", err))
	}
	if err = v.Unmarshal(&g.Config); err != nil {
		// 将配置文件反序列化到 Config 结构体
		panic(fmt.Errorf("unmarshal config failed, err: %v", err))
	}
}
