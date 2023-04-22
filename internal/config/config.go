package config

import (
	"github.com/behrouz-rfa/mongo-specification/pkg/database/factory"
	"github.com/spf13/viper"
	"gopkg.in/jeevatkm/go-model.v1"
	"log"
	"os"
)

type Config struct {
	Host     string
	Port     int
	GinMode  string
	Mongo    Mongo
	Logger   LoggerConfig
	Services servicesAddr
	Nats     Nats
}
type Mongo struct {
	Host      string
	Port      int
	Username  string
	Password  string
	Name      string
	Driver    string
	SSL       string
	Timeout   int
	Clustered bool
}

func (p Mongo) ToPgConfig() factory.MongoConfig {
	f := factory.MongoConfig{}
	errors := model.Copy(&f, p)
	if len(errors) > 0 {
		panic(errors[0])
	}
	return f
}

type Nats struct {
	Url   string
	Port  int
	Topic Topics
}

type Topics struct {
	Read  string
	Write string
}
type servicesAddr struct {
	Workflow string
}
type LoggerConfig struct {
	//Console *logger.ConsoleConfig
}

func LoadConfig() *Config {
	path := os.Getenv("CONFIG_PATH")
	if path == "" {
		path = "configs/"
	}

	log.Println("Loading", "path", path)

	v := viper.New()
	v.AddConfigPath(path)
	v.SetConfigName("config")

	if err := v.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	c := &Config{}

	if err := v.Unmarshal(c); err != nil {
		log.Fatal(err)
	}

	log.Println("%+v", c)

	return c
}
