package config

import (
	"strings"
	"sync"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	configInst *viper.Viper
	envInst    *viper.Viper

	configOnce sync.Once
	envOnce    sync.Once
)

func ConForge() *viper.Viper {
	configOnce.Do(func() {
		configInst = viper.New()
		configInst.SetConfigName("config")
		configInst.SetConfigType("json")
		configInst.AddConfigPath(".")
		configInst.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
		configInst.AutomaticEnv()

		if err := configInst.ReadInConfig(); err != nil {
			log.Error("Could not find Config file, error " + err.Error())
		}

	})

	return configInst
}

func EnvForge() *viper.Viper {
	envOnce.Do(func() {
		envInst = viper.New()
		envInst.SetConfigName("env")
		envInst.SetConfigType("yml")
		envInst.AddConfigPath(".")
		envInst.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
		envInst.AutomaticEnv()

		err := envInst.ReadInConfig()
		if err != nil {
			log.Error(err)
		}

		if err := envInst.ReadInConfig(); err != nil {
			log.Error("Could not find env configuration file, error " + err.Error())
		}

	})
	return envInst
}
