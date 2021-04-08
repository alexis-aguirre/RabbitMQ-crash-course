package util

import (
	"github.com/spf13/viper"
)

var globalConfig Config

type Config struct {
	RabbitConfig rabbitConfig
	QueueConfig  queueConfig
}

type rabbitConfig struct {
	PORT           string `mapstructure:"QUEUE_PORT"`
	HOST           string `mapstructure:"QUEUE_HOST"`
	RabbitUser     string `mapstructure:"QUEUE_USER"`
	RabbitPassword string `mapstructure:"QUEUE_PASSWORD"`
}

type queueConfig struct {
	ExchangeName string `mapstructure:"EXCHANGE_NAME"`
	QueueName    string `mapstructure:"QUEUE_NAME"`
	RoutingKey   string `mapstructure:"QUEUE_ROUTING_KEY"`
}

func LoadConfig(path string) (Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	queueConf := queueConfig{}
	rabbitConf := rabbitConfig{}

	err := viper.ReadInConfig()
	if err != nil {
		return globalConfig, err
	}

	err = viper.Unmarshal(&queueConf)
	if err != nil {
		return globalConfig, err
	}

	err = viper.Unmarshal(&rabbitConf)
	if err != nil {
		return globalConfig, err
	}

	globalConfig.QueueConfig = queueConf
	globalConfig.RabbitConfig = rabbitConf
	return globalConfig, nil

}

func GetConfig() Config {
	return globalConfig
}
