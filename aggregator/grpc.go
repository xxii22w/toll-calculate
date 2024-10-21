package main

import (
	"context"

	"tolling/types"
)

type GRPCDistanceAggregatorServer struct {
	types.UnimplementedDistanceAggregatorServer
	svc Aggregator
}

func NewGRPCDistanceAggregatorServer(svc Aggregator) *GRPCDistanceAggregatorServer {
	return &GRPCDistanceAggregatorServer{
		svc: svc,
	}
}

// transport layer
// JSON -> types.Distance -> all done (same type)
// GRPC -> types.AggregateRequest -> type.Distance
// Webpack => types.WEBpack -> types.Distance

// business layer -> business layer type (main type everyone needs to convert to)

func (s *GRPCDistanceAggregatorServer) AggregateDistance(ctx context.Context, req *types.AggregatorDistanceRequest) (*types.None, error) {
	distance := types.Distance{
		OBUID: int(req.ObuID),
		Value: req.Value,
		Unix:  req.Unix,
	}
	return &types.None{}, s.svc.AggregateDistance(distance)
}

func (s *GRPCDistanceAggregatorServer) GetInvoice(ctx context.Context, req *types.GetInvoiceRequest) (*types.GetInvoiceResponse, error) {
	inv, err := s.svc.Invoice(int(req.ObuID))
	if err != nil {
		return nil, err
	}
	return &types.GetInvoiceResponse{
		ObuID:         int64(inv.OBUID),
		InvoiceAmount: inv.InvoiceAmount,
		TotalDistance: inv.TotalDistance,
	}, nil
}
