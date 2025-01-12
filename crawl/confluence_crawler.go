package crawl

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type ConfluenceCrawler struct {
	BaseURL  string
	Username string
	APIToken string
}

func NewConfluenceCrawler(baseURL, username, apiToken string) *ConfluenceCrawler {
	return &ConfluenceCrawler{BaseURL: baseURL, Username: username, APIToken: apiToken}
}

func (cc *ConfluenceCrawler) Crawl() ([]byte, error) {
	url := fmt.Sprintf("%s/rest/api/content", cc.BaseURL)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(cc.Username, cc.APIToken)
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
