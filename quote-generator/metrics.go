package main

import (
	"context"
	"time"

	"github.com/go-kit/kit/metrics"
)

type instrumentingMiddleware struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	next           SaleGeneratorService
}

func (mw instrumentingMiddleware) Sale(ctx context.Context) (result Sale, err error) {
	defer func(begin time.Time) {
		mw.requestCount.With("method", "Sale", "error", err.Error()).Add(1)
		// mw.requestLatency.With("method", "Sale", "error", err.Error()).Observe(time.Since(begin).Seconds())
	}(time.Now())

	result, err = mw.next.Sale(ctx)
	return
}
