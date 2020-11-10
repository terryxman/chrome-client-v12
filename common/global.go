package Common

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/spf13/viper"
	"time"
)

var Config = viper.New()

var Json = jsoniter.ConfigCompatibleWithStandardLibrary

var TimeLocal, _ = time.LoadLocation("Asia/Shanghai")

var ConsoleConfig = viper.New()
