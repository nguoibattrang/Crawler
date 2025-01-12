package config

import (
	"fmt"
	"time"

	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Crawlers []Crawler `yaml:"crawlers"`
	Kafka    Kafka     `yaml:"kafka"`
}

type Crawler interface{}

type JiraCrawler struct {
	Type           string        `yaml:"type"`
	BaseURL        string        `yaml:"base_url"`
	APIToken       string        `yaml:"api_token"`
	Username       string        `yaml:"username"`
	Interval       time.Duration `yaml:"interval"`
	OtherJiraField string        `yaml:"other_jira_field"`
}

// ConfluenceCrawler struct
type ConfluenceCrawler struct {
	Type     string        `yaml:"type"`
	BaseURL  string        `yaml:"base_url"`
	APIToken string        `yaml:"api_token"`
	Username string        `yaml:"username"`
	Interval time.Duration `yaml:"interval"`
}

// Kafka struct
type Kafka struct {
	Brokers []string          `yaml:"brokers"`
	Topics  map[string]string `yaml:"topics"`
}

func (c *Config) UnmarshalYAML(value *yaml.Node) error {
	type Alias Config // Create an alias to avoid recursion
	var temp struct {
		Crawlers []map[string]interface{} `yaml:"crawlers"`
		Kafka    Kafka                    `yaml:"kafka"`
	}
	if err := value.Decode(&temp); err != nil {
		return err
	}

	c.Kafka = temp.Kafka
	for _, crawler := range temp.Crawlers {
		switch crawler["type"] {
		case "jira":
			var jira JiraCrawler
			if err := decodeMapToStruct(crawler, &jira); err != nil {
				return err
			}
			c.Crawlers = append(c.Crawlers, &jira)
		case "confluence":
			var confluence ConfluenceCrawler
			if err := decodeMapToStruct(crawler, &confluence); err != nil {
				return err
			}
			c.Crawlers = append(c.Crawlers, &confluence)
		default:
			return fmt.Errorf("unsupported crawler type: %s", crawler["type"])
		}
	}
	return nil
}

// Helper to decode map into a struct
func decodeMapToStruct(data map[string]interface{}, target interface{}) error {
	yamlBytes, err := yaml.Marshal(data)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(yamlBytes, target)
}

func Parse(filePath string) (*Config, error) {

	// Open YAML file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Decode YAML into struct
	var config Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
