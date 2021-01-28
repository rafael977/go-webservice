package server

import (
	"context"
	"currency/protos"
	"github.com/hashicorp/go-hclog"
)

type Currency struct {
	log hclog.Logger
	grpc.UnimplementedCurrencyServer
}

func NewCurrency(l hclog.Logger) *Currency {
	return &Currency{log: l}
}

func (c *Currency) GetRate(ctx context.Context, rr *grpc.RateRequest) (*grpc.RateResponse, error) {
	c.log.Info("GetRate", "base", rr.GetBase(), "target", rr.GetTarget())

	return &grpc.RateResponse{Rate: 0.5}, nil
}
