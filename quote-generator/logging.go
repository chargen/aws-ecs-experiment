package main

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"
)

type loggingMiddleware struct {
	logger log.Logger
	next   SaleGeneratorService
}

func (mw loggingMiddleware) Sale(ctx context.Context) (result Sale, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "Sale",
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	result, err = mw.next.Sale(ctx)
	return
}
