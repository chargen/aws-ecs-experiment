package main

import (
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	logger := log.NewJSONLogger(os.Stderr)

	var service SaleGeneratorService
	service = saleGeneratorService{}
	service = loggingMiddleware{logger, service}

	saleHandler := httptransport.NewServer(
		makeSaleEndpoint(service),
		decodeSaleRequest,
		encodeResponse,
	)

	http.Handle("/sale", saleHandler)
	http.ListenAndServe(":8080", nil)
}
