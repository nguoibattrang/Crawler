package crawl

import (
	"data_crawler/config"
	"fmt"
	"go.uber.org/zap"
	"os"
	"path/filepath"
)

type FileCrawler struct {
	Path   string
	logger *zap.Logger
}

func NewFileCrawler(cfg *config.CrawlerConfig, logger *zap.Logger) *FileCrawler {
	return &FileCrawler{Path: cfg.Path, logger: logger}
}

func (inst *FileCrawler) Crawl(chanMsg chan<- string) {
	// Read all files in the directory
	files, err := os.ReadDir(inst.Path)
	if err != nil {
		inst.logger.Error("Failed to read directory", zap.Error(err))
	}

	for _, file := range files {
		if file.IsDir() {
			// Skip directories
			inst.logger.Info("Skipping directory", zap.String("dir", file.Name()))
			continue
		}

		filePath := filepath.Join(inst.Path, file.Name())

		content, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Printf("Error reading file %s: %v\n", filePath, err)
			continue
		}
		chanMsg <- string(content)
	}
}
