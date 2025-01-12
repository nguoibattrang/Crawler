package crawl

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type JiraCrawler struct {
	BaseURL  string
	Username string
	APIToken string
}

func NewJiraCrawler(baseURL, username, apiToken string) *JiraCrawler {
	return &JiraCrawler{BaseURL: baseURL, Username: username, APIToken: apiToken}
}

func (jc *JiraCrawler) Crawl() ([]byte, error) {
	url := fmt.Sprintf("%s/rest/api/2/search", jc.BaseURL)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(jc.Username, jc.APIToken)
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
