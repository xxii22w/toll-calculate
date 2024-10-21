package client

import (
	"context"
	"tolling/types"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCClient struct {
	Endpoint string
	client   types.DistanceAggregatorClient
}

func NewGRPCClient(endpoint string) (*GRPCClient, error) {
	conn, err := grpc.Dial(":3001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	c := types.NewDistanceAggregatorClient(conn)
	return &GRPCClient{
		Endpoint: endpoint,
		client:   c,
	}, nil
}

func (c *GRPCClient) AggregateDistance(ctx context.Context, req *types.AggregatorDistanceRequest) error {
	_, err := c.client.AggregateDistance(ctx, req)
	if err != nil {
		return err
	}
	return nil
}

func (gc *GRPCClient) GetInvoice(ctx context.Context, req *types.GetInvoiceRequest) (*types.Invoice, error) {
	resp, err := gc.client.GetInvoice(ctx, req)
	if err != nil {
		return nil, err
	}
	return &types.Invoice{
		OBUID:         int(resp.ObuID),
		TotalDistance: resp.TotalDistance,
		InvoiceAmount: resp.InvoiceAmount,
	}, nil
}
