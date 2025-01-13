package main

import (
	"fmt"
	"github.com/nguoibattrang/crawler/config"
	"github.com/nguoibattrang/crawler/crawl"
	"github.com/nguoibattrang/crawler/logger"
	"github.com/nguoibattrang/crawler/output"
	"go.uber.org/zap"
	"os"
	"path/filepath"
)

func main() {
	// Load configuration
	serviceCfg, err := config.LoadConfig(filepath.Join(os.Getenv("CONFIG_PATH"), "config.yaml"))
	if err != nil {
		fmt.Printf("config.LoadConfig fail to load config %v", err)
		os.Exit(1)
	}

	log, err := logger.InitLogger(serviceCfg.Logger.Mode)
	if err != nil {
		fmt.Printf("logger.InitLogger failed to init logger %v", err)
		os.Exit(1)
	}

	producer, err := output.NewKafkaProducer(serviceCfg.Kafka.Brokers, serviceCfg.Kafka.Topic, log)
	if err != nil {
		log.Fatal("Failed to init kafka producer", zap.Error(err))
		os.Exit(1)
	}

	crawlers := crawl.CreateCrawlers(serviceCfg.Crawlers, log)
	mChan := make(chan string)
	for _, c := range crawlers {
		go c.Crawl(mChan)
	}

	producer.Produce(mChan)
}
