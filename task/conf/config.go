package conf

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"task/model"
)

type RabbitMQ struct {
	Rabbitmq         string
	RabbitmqUser     string
	RabbitmqPassword string
	RabbitmqHost     string
	RabbitmqPort     string
}

type Sql struct {
	Db       string
	DbHost   string
	DbPort   int
	DbUser   string
	Password string
	DbName   string
}

type Config struct {
	Sql Sql
	MQ  RabbitMQ
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

func Init() {
	viper.SetConfigType("yaml")
	viper.SetConfigFile("./conf/config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println(err)
	}
	var myConfig Config
	err = viper.Unmarshal(&myConfig)
	if err != nil {
		log.Println(err)
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?"+
		"charset=utf8mb4&parseTime=True&loc=Local", myConfig.Sql.DbUser,
		myConfig.Sql.Password, myConfig.Sql.DbHost, myConfig.Sql.DbPort,
		myConfig.Sql.DbName)
	err = model.Database(dsn)
	if err != nil {
		log.Println(err)
	}
}
