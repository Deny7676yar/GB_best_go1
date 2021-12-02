package crawlerer

import (
	"context"
	"sync"
)

type Crawler interface {
	Scan(ctx context.Context, wg *sync.WaitGroup, url string, depth int)
	ChanResult() <-chan CrawlResult
	ToChanResult(CrawlResult)
}

type Requester interface {
	Get(ctx context.Context, url string) (Page, error)
}

type Page interface {
	GetTitle() string
	GetLinks() []string
}

type CrawlResult struct {
	Err   error
	Info  string
	Title string
	URL   string
}
