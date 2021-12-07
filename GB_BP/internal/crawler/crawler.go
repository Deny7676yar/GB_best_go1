package crawler

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	cfg "github.com/Deny7676yar/Go_level2/GB_BP/internal/config"
	cra "github.com/Deny7676yar/Go_level2/GB_BP/internal/controllers/crawlerer"
	log "github.com/sirupsen/logrus"
)

//Crawler - интерфейс (контракт) краулера
type Crawler interface {
	Scan(ctx context.Context, wg *sync.WaitGroup, url string, depth int)
	ChanResult() <-chan cra.CrawlResult
}

type crawler struct {
	r           cra.Requester
	res         chan cra.CrawlResult
	visited     map[string]struct{}
	mu          sync.RWMutex
	searchDepth int
}

func NewCrawler(r cra.Requester) *crawler {
	return &crawler{
		r:       r,
		res:     make(chan cra.CrawlResult),
		visited: make(map[string]struct{}),
		mu:      sync.RWMutex{},
	}
}

func (c *crawler) Scan(ctx context.Context, wg *sync.WaitGroup, url string, depth int) {
	//if depth <= 0 { //Проверяем то, что есть запас по глубине
	//	return
	//}

	defer wg.Done()
	if depth >= c.searchDepth {
		return
	}
	c.mu.RLock()
	_, ok := c.visited[url] //Проверяем, что мы ещё не смотрели эту страницу
	c.mu.RUnlock()
	if ok {
		return
	}
	select {
	case <-ctx.Done(): //Если контекст завершен - прекращаем выполнение
		return
	default:
		page, err := c.r.Get(ctx, url) //Запрашиваем страницу через Requester
		if err != nil {
			c.res <- cra.CrawlResult{Err: err} //Записываем ошибку в канал
			return
		}
		c.mu.Lock()
		c.visited[url] = struct{}{} //Помечаем страницу просмотренной
		c.mu.Unlock()
		c.res <- cra.CrawlResult{ //Отправляем результаты в канал
			Title: page.GetTitle(),
			URL:   url,
		}
		for _, link := range page.GetLinks() {
			go c.Scan(ctx, wg, link, depth-1) //На все полученные ссылки запускаем новую рутину сборки
		}
	}
}

func (c *crawler) ChanResult() <-chan cra.CrawlResult {
	return c.res
}

func SearchDepthCrawler(maxDepth int) *crawler {
	return &crawler{
		searchDepth: maxDepth,
	}
}

func SigDepth(ctx context.Context, c *crawler, d int) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGUSR1)
	for {
		select {
		case <-ctx.Done():
			return
		case <-sigChan:
			log.WithFields(log.Fields{
				"SIGUSR1": <-sigChan,
			}).Info("Depth += 2")
			c.InitDepth(d)
		}
	}
}

func (c *crawler) InitDepth(dep int) {
	c.mu.Lock()
	c.searchDepth += dep
	c.mu.Unlock()
}

func ProcessResult(ctx context.Context, cancel func(), cr Crawler, cfg cfg.Config) {
	var maxResult, maxErrors = cfg.MaxResults, cfg.MaxErrors
	for {
		select {
		case <-ctx.Done():
			return
		case msg := <-cr.ChanResult():
			if msg.Err != nil {
				maxErrors--
				log.Printf("crawler result return err: %s\n", msg.Err.Error())
				if maxErrors <= 0 {
					cancel()
					return
				}
			} else {
				maxResult--
				log.Printf("crawler result: [url: %s] Title: %s\n", msg.URL, msg.Title)
				if maxResult <= 0 {
					cancel()
					return
				}
			}
		}
	}
}
