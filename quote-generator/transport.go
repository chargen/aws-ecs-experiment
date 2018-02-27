package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

type saleRequest struct{}

type saleResponse struct {
	V   Sale   `json:"v"`
	Err string `json:"err,omitempty"`
}

func makeSaleEndpoint(svc SaleGeneratorService) endpoint.Endpoint {
	return func(ctx context.Context, _ interface{}) (interface{}, error) {
		v, err := svc.Sale(ctx)
		if err != nil {
			return saleResponse{v, err.Error()}, nil
		}
		return saleResponse{v, ""}, nil
	}
}

func decodeSaleRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request saleRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
