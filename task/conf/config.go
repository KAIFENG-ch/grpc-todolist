package conf

import (
	"github.com/spf13/viper"
	"log"
)

type RabbitMQ struct {
	Rabbitmq string
	RabbitmqUser string
	RabbitmqPassword string
	RabbitmqHost string
	RabbitmqPort string
}

func LoadRabbitMQ() string {
	viper.SetConfigType("yaml")
	viper.SetConfigFile("./conf/config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println(err)
	}
	var rabbitConf RabbitMQ
	err = viper.Unmarshal(&rabbitConf)
	if err != nil {
		log.Println(err)
	}
	RabbitMQPath := rabbitConf.Rabbitmq + "://" + rabbitConf.RabbitmqUser + ":" +
		rabbitConf.RabbitmqPassword + "@" + rabbitConf.RabbitmqHost + ":" +
		rabbitConf.RabbitmqPort + "/"
	return RabbitMQPath
}