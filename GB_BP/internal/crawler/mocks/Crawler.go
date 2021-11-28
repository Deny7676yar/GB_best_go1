// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"
	"github.com/Deny7676yar/Go_level2/GB_BP/internal/crawler"
)
import mock "github.com/stretchr/testify/mock"
import sync "sync"

// Crawler is an autogenerated mock type for the Crawler type
type Crawler struct {
	mock.Mock
}

// ChanResult provides a mock function with given fields:
func (_m *Crawler) ChanResult() <-chan crawler.CrawlResult {
	ret := _m.Called()

	var r0 <-chan crawler.CrawlResult
	if rf, ok := ret.Get(0).(func() <-chan crawler.CrawlResult); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan crawler.CrawlResult)
		}
	}

	return r0
}

// Scan provides a mock function with given fields: ctx, wg, url, depth
func (_m *Crawler) Scan(ctx context.Context, wg *sync.WaitGroup, url string, depth int) {
	_m.Called(ctx, wg, url, depth)
}
