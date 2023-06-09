package config

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
	"sync"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/kr/pretty"
	"github.com/spf13/viper"
)

const (
	ProjectName = "hertz-scaffold"
	EnvDev      = "dev"
	EnvTest     = "test"
	EnvOnline   = "online"
)

var (
	env  string
	conf *AutoGenerated
	once sync.Once
)

func init() {
	once.Do(Init)
}

// GetConf gets configuration instance
func GetConf() AutoGenerated {
	return *conf
}

func GetEnv() string {
	return env
}

// Init reads config if set.
// Set > Flag > Env > Config File > key/Value > Defaults
func Init() {
	viper.SetDefault("env", EnvDev)

	env = viper.GetString("env")
	if env != EnvTest && env != EnvDev && env != EnvOnline {
		fmt.Println("env is incorrect, use default env: dev")
		os.Exit(1)
	}

	// 读取环境变量
	// viper.SetEnvPrefix(strings.ToUpper(ProjectName))
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	viper.SetConfigName(env)   // 配置文件名，不需要后缀名
	viper.SetConfigType("yml") // 配置文件格式
	viper.AddConfigPath(getConfigRoot())

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Printf("config not found！%v\n", err)
		} else {
			fmt.Printf("config read failed! %v\n", err)
		}
		os.Exit(1)
	}
	if err := viper.Unmarshal(&conf); err != nil {
		fmt.Printf("unmarshal config file failed, %v\n", err)
		os.Exit(1)
	}
	hlog.Info(pretty.Sprintf("Using config file[%s]\n", viper.ConfigFileUsed()))
	hlog.Info(pretty.Sprint(conf))
}

func getConfigRoot() string {
	_, current, _, _ := runtime.Caller(0)
	return path.Dir(current)
}
