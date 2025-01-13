package crawl

type Crawler interface {
	Crawl(chan<- Data)
}
