package conf

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"user/model"
)

type Config struct {
	Sql Sql
}

type Sql struct {
	Db       string
	DbHost   string
	DbPort   int
	DbUser   string
	Password string
	DbName   string
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
