package crawl

type Crawler interface {
	Crawl() ([]byte, error) // Crawls data and returns the result as JSON or raw bytes
}
