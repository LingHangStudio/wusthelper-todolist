package conf

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/spf13/viper"
	"time"
	"wusthelper-todolist-service/library/database"
)

const (
	DevEnv  = "dev"
	ProdEnv = "prod"
)

var (
	Conf = &Config{}
)

type Config struct {
	Server   ServerConf
	Database database.Config
}

type ServerConf struct {
	Env            string
	Port           int
	Address        string
	BaseUrl        string
	TokenSecret    string
	PasswordKey    string
	TokenTimeout   time.Duration
	LogLocation    string
	VersionLogFile string
}

func Init() (err error) {
	viper.AddConfigPath(".")
	viper.AddConfigPath("./conf")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("/etc/wusthelper-mp")
	viper.AddConfigPath("$HOME/.wusthelper-mp")

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(Conf)
	if err != nil {
		return
	}

	if Conf.Server.Env == DevEnv {
		jsonByte, _ := jsoniter.Marshal(Conf)
		fmt.Println(string(jsonByte))
	}

	return
}
