package crawl

import (
	"github.com/nguoibattrang/crawler/config"
	"go.uber.org/zap"
)

func Create(cfg *config.CrawlerConfig, log *zap.Logger) Crawler {
	switch cfg.Type {
	case "jira":
		return NewJiraCrawler(cfg, log)
	case "confluence":
		return NewConfluenceCrawler(cfg, log)
	case "file":
		return NewFileCrawler(cfg, log)
	}
	return nil
}

func CreateCrawlers(list []*config.CrawlerConfig, log *zap.Logger) []Crawler {
	var crawlers []Crawler
	for _, cfg := range list {
		c := Create(cfg, log)
		if c != nil {
			crawlers = append(crawlers, c)
		}
	}
	return crawlers
}
