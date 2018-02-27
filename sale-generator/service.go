package main

import (
	"context"

	"github.com/Pallinder/go-randomdata"
)

type Feature struct {
	Name  string `json:"feature"`
	Value string `json:"value"`
}

type Sale struct {
	ProductGroup string    `json:"product_group"`
	Product      string    `json:"product"`
	Brand        string    `json:"brand"`
	Price        float64   `json:"price"`
	Date         string    `json:"time"`
	Features     []Feature `json:"features"`
}

type SaleGeneratorService interface {
	Sale(context.Context) (Sale, error)
}

type saleGeneratorService struct{}

func (saleGeneratorService) Sale(_ context.Context) (Sale, error) {
	sale := Sale{
		ProductGroup: randomdata.StringSample("phone", "television", "computer"),
		Product:      randomdata.SillyName(),
		Brand:        randomdata.SillyName(),
		Price:        randomdata.Decimal(2000),
		Date:         randomdata.FullDate(),
		Features: []Feature{
			{Name: "os", Value: randomdata.StringSample("Linux", "Unix", "Windows")},
			{Name: "storage_gb", Value: randomdata.StringSample("16", "32", "64")},
			{Name: "wifi", Value: randomdata.StringSample("yes", "no")},
			{Name: "bluetooth", Value: randomdata.StringSample("yes", "no")},
			{Name: "nfc", Value: randomdata.StringSample("yes", "no")},
		},
	}

	return sale, nil
}

type ServiceMiddleware func(SaleGeneratorService) SaleGeneratorService
