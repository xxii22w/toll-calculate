package client

import (
	"context"

	"tolling/types"
)

type Client interface {
	AggregateDistance(context.Context, *types.AggregatorDistanceRequest) error
	GetInvoice(context.Context, *types.GetInvoiceRequest) (*types.Invoice, error)
}
