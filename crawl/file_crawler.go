package crawl

import (
	"os"
	"path/filepath"

	"github.com/nguoibattrang/crawler/config"
	"go.uber.org/zap"
)

type FileCrawler struct {
	Path   string
	logger *zap.Logger
	Site   string
}

func NewFileCrawler(cfg *config.CrawlerConfig, logger *zap.Logger) *FileCrawler {
	return &FileCrawler{Path: cfg.Path, logger: logger, Site: cfg.Site}
}

func (inst *FileCrawler) Crawl(chanMsg chan<- Data) {
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
			inst.logger.Error("Error reading file", zap.String("filepath", filePath), zap.Error(err))
			continue
		}

		chanMsg <- Data{
			Type:    inst.Site,
			Content: string(content),
		}
	}
}
