package config

import (
	"sync"
)

// Config is a struct that contains configuration variables
type Config struct {
	App      *App
	Server   *ServerList
	Database *DatabaseList
}

var Instance *Config

var configInstance *Config
var once sync.Once

func Get() *Config {
	once.Do(func() {
		configInstance = LoadConfig()
	})

	return configInstance
}
