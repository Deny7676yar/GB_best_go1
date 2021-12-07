package services

import (
	"context"
	"github.com/Deny7676yar/Go_level2/GB_BP/internal/page"
	"net/http"
	"time"

	cra "github.com/Deny7676yar/Go_level2/GB_BP/internal/controllers/crawlerer"
	log "github.com/sirupsen/logrus"
)

type Requester interface {
	Get(ctx context.Context, url string) (cra.Page, error)
}

type requester struct {
	timeout time.Duration
}

func NewRequester(timeout time.Duration) requester {
	return requester{timeout: timeout}
}

func (r requester) Get(ctx context.Context, url string) (cra.Page, error) {
	select {
	case <-ctx.Done():
		return nil, nil
	default:
		cl := &http.Client{
			Timeout: r.timeout,
		}
		req, err := http.NewRequest("GET", url, nil) //nolint
		if err != nil {
			return nil, err
		}
		body, err := cl.Do(req)
		if err != nil {
			log.WithFields(log.Fields{
				"Get body:": err,
			}).Errorf("Do Not body")
			return nil, err
		}
		defer body.Body.Close()

		newPage, err := page.NewPage(body.Body)
		if err != nil {
			log.WithFields(log.Fields{
				"NewPage:": err,
			}).Panicf("No Page")
			return nil, err
		}
		return newPage, nil
	}
}
