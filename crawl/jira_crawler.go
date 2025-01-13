package crawl

import (
	"fmt"
	"github.com/nguoibattrang/crawler/config"
	"go.uber.org/zap"
	"io"
	"net/http"
)

type JiraCrawler struct {
	url    string
	token  string
	logger *zap.Logger
}

func NewJiraCrawler(cfg *config.CrawlerConfig, logger *zap.Logger) *JiraCrawler {
	return &JiraCrawler{url: cfg.Path, token: cfg.APIToken, logger: logger}
}

func (inst *JiraCrawler) Crawl(chanMsg chan<- string) {
	url := fmt.Sprintf("%s/rest/api/2/search", inst.url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		inst.logger.Error("Failed to create request", zap.Error(err))
		return
	}
	req.SetBasicAuth("", inst.token)
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
