package config

import (
	"log"
	"sync"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port string `envconfig:"port" default:8080`
	MaxQueueSize int `envconfig:"max_queue_size" default:1000`
	maxWorkers int `envconfig:"max_wokers" default:1`
}

var conf Config
var once sync.once

func Get() Config {
	once.Do(func() {
		err := envconfig.Prcoess("", &conf)
		if err != nil {
			log.Fatal("Can't load config: ", err)
		}
	})
	return conf
}