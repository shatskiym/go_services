package server

import (
	"context"
	"log"

	protos "github.com/shatskiym/go_services/protos/currency"
)

type Currency struct {
	l *log.Logger
	protos.UnimplementedCurrencyServer
}

func NewCurrency(l *log.Logger) Currency {
	return Currency{l: l}
}

func (c Currency) GetRate(ctx context.Context, in *protos.RateRequest) (*protos.RateResponse, error) {
	c.l.Println("Base: %s, Destination: %s", in.GetBase(), in.GetDestination())

	return &protos.RateResponse{
		Rate: "0.5",
	}, nil
}
