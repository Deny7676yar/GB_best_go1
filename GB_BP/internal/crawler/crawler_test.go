package crawler

import (
	"github.com/Deny7676yar/Go_level2/GB_BP/internal/services"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewCrawler(t *testing.T) {
	r := services.NewRequester(10 * time.Second)
	cr := NewCrawler(r)
	assert.NotNil(t, cr, "Error on create new Crawler, get nil.")
}

//func TestCrawler_Scan(t *testing.T) {
//	wg := &sync.WaitGroup{}
//	wg.Add(1)
//	wg.Done()
//
//	r := &mocks.Requester{}
//	pg := &mocks2.Page{}
//	cr := NewCrawler(r)
//
//	ctx := context.Background()
//	var depth int64 = 3
//	var start int64 = 1
//	url := "https://testStartURL.com"
//
//	exVisitedURLCount := 4
//	exTitle := "Test title"
//	exLinks := []string{"https://www.yandex.ru", "https://www.google.com", "https://www.mozilla.org"}
//
//	pg.On("GetTitle").Return(exTitle)
//	pg.On("GetLinks").Return(exLinks)
//
//	r.On("Get", ctx, url).Return(pg, nil)
//	r.On("Get", ctx, exLinks[0]).Return(pg, nil)
//	r.On("Get", ctx, exLinks[1]).Return(pg, nil)
//	r.On("Get", ctx, exLinks[2]).Return(pg, nil)
//
//	go cr.Scan(ctx, wg, url, &depth, start)
//
//	var maxResult, maxErrors = 10, 5
//	doFor := true
//	for doFor {
//		select {
//		case msg := <-cr.ChanResult():
//			if msg.Err != nil {
//				maxErrors--
//				if maxErrors <= 0 {
//					doFor = false
//				}
//			} else if len(msg.Info) > 0 {
//				doFor = false
//			} else {
//				maxResult--
//				if maxResult <= 0 {
//					doFor = false
//				}
//			}
//		}
//	}
//
//	assert.Equal(t, exVisitedURLCount, len(cr.visited), "Not equal.\n Expected: %v \n Got: %v \n", exVisitedURLCount, len(cr.visited))
//}
