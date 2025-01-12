package main

import (
	"data_crawler/crawl"
	"log"
	"time"

	"data_crawler/config"
	"data_crawler/output"
	"data_crawler/scheduler"
)

func main() {
	// Load configuration
	cfg, err := config.Parse("config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	crawlers := crawl.CreateCrawler(cfg.Crawlers)

	for _, c := range crawlers {
		p := output.NewKafkaProducer(nil, "")
		go scheduler.StartScheduler(1*time.Hour, func() {
			data, err := c.Crawl()
			if err != nil {
				log.Printf("Failed to crawl Confluence: %v", err)
				return
			}
			if err := p.Produce(data); err != nil {
				log.Printf("Failed to publish Confluence data: %v", err)
			}
		})
	}

	select {}
}

func parseDuration(interval string) time.Duration {
	d, err := time.ParseDuration(interval)
	if err != nil {
		log.Fatalf("Invalid duration format: %v", err)
	}
	return d
}
