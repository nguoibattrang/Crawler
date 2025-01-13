package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type ServiceConfig struct {
	Crawlers []*CrawlerConfig `yaml:"crawlers"`
	Kafka    *KafkaConfig     `yaml:"kafka"`
	Logger   *LogConfig       `yaml:"logger"`
}

type CrawlerConfig struct {
	Type     string `yaml:"type"`
	Path     string `yaml:"path"`
	APIToken string `yaml:"api_token"`
	Site     string `yaml:"site"`
}

type LogConfig struct {
	Mode string `yaml:"mode"`
}

type KafkaConfig struct {
	Brokers []string `yaml:"brokers"`
	Topic   string   `yaml:"topic"`
}

func LoadConfig(filename string) (*ServiceConfig, error) {
	// Initialize viper
	v := viper.New()
	v.SetConfigFile(filename)
	v.SetConfigType("yaml")

	// Set defaults if needed
	v.SetDefault("logger.mode", "production")

	// Read in the configuration file
	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	// Unmarshal into the ServiceConfig struct
	var config ServiceConfig
	if err := v.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &config, nil
}
