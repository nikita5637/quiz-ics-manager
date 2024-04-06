package config

import (
	"fmt"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	defaultLogLevel = zap.NewAtomicLevelAt(zap.InfoLevel)
)

func initLoggerConfigureParams() {
	_ = viper.BindEnv("log.level")
	_ = viper.BindEnv("log.module_name")
	_ = viper.BindEnv("log.elastic.address")
	_ = viper.BindEnv("log.elastic.enabled")
	_ = viper.BindEnv("log.elastic.index")
	_ = viper.BindEnv("log.elastic.port")
}

// GetElasticAddress ...
func GetElasticAddress() string {
	return fmt.Sprintf("http://%s:%d",
		viper.GetString("log.elastic.address"),
		viper.GetUint32("log.elastic.port"),
	)
}

// GetLogLevel ...
func GetLogLevel() zap.AtomicLevel {
	level, err := zap.ParseAtomicLevel(viper.GetString("log.level"))
	if err != nil {
		level = defaultLogLevel
	}

	return level
}
