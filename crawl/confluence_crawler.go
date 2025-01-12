package crawl

import (
	"data_crawler/config"
	"fmt"
	"go.uber.org/zap"
	"io"
	"net/http"
)

type ConfluenceCrawler struct {
	BaseURL  string
	APIToken string
	logger   *zap.Logger
}

func NewConfluenceCrawler(cfg *config.CrawlerConfig, logger *zap.Logger) *ConfluenceCrawler {
	return &ConfluenceCrawler{BaseURL: cfg.Domain, APIToken: cfg.APIToken, logger: logger}
}

func (inst *ConfluenceCrawler) Crawl(chanMsg chan<- string) {
	url := fmt.Sprintf("%s/rest/api/2/search", inst.BaseURL)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		inst.logger.Error("Failed to create request", zap.Error(err))
		return
	}
	req.SetBasicAuth("", inst.APIToken)
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		inst.logger.Error("Failed to send request", zap.Error(err))
		return
	}
	defer resp.Body.Close()

	msg, err := io.ReadAll(resp.Body)
	if err != nil {
		inst.logger.Error("Failed to read response", zap.Error(err))
		return
	}
	chanMsg <- string(msg)
}
