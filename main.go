package main

import (
	"data_crawler/config"
	"data_crawler/crawl"
	"data_crawler/logger"
	"data_crawler/output"
	"fmt"
	"os"
)

func main() {
	// Load configuration
	serviceCfg, err := config.LoadConfig("C:\\Users\\tient\\Desktop\\Crawler\\app.yml")
	if err != nil {
		fmt.Printf("config.LoadConfig fail to load config %v", err)
		os.Exit(1)
	}

	log, err := logger.InitLogger(serviceCfg.Logger.Mode)
	if err != nil {
		fmt.Printf("logger.InitLogger failed to init logger %v", err)
		os.Exit(1)
	}

	crawlers := crawl.CreateCrawlers(serviceCfg.Crawlers, log)
	producer := output.NewKafkaProducer(serviceCfg.Kafka.Brokers, serviceCfg.Kafka.Topics, log)

	mChan := make(chan string)
	for _, c := range crawlers {
		go c.Crawl(mChan)
	}

	producer.Produce(mChan)
	log.Info("crawler exited")
}
