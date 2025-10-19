package config

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func LoadConfig() *Config {
	_ = godotenv.Load()

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Errorf("failed to read config: %w", err)
		panic("cannot read config file")
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		fmt.Errorf("failed to unmarshal config: %w", err)
		panic("cannot read config file")
	}

	fmt.Printf("Configuration loaded: %+v\n", cfg)

	return &cfg
}
