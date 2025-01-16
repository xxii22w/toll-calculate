package aggendpoint

import (
	"context"
	"time"
	"tolling/go-kit-example/aggsrv/aggservice"
	"tolling/types"

	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/metrics/prometheus"
	"github.com/go-kit/kit/ratelimit"
	"github.com/go-kit/log"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/sony/gobreaker"
	"golang.org/x/time/rate"
)

type Set struct {
	AggregateEndpoint endpoint.Endpoint
	CalculateEndpoint endpoint.Endpoint
}

func New(svc aggservice.Service, logger log.Logger) Set {
	duration := prometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "toll_calculator",
		Subsystem: "aggregation",
		Name:      "request_duration_seconds",
		Help:      "Request duration measured in seconds",
	}, []string{"method", "success"})

	var aggregateEndpoint = MakeAggregateEndpoint(svc)
	{
		aggregateEndpoint = ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Every(time.Second), 1))(aggregateEndpoint)
		aggregateEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(aggregateEndpoint)
		aggregateEndpoint = LoggingMiddleware(log.With(logger, "method", "Aggregate"))(aggregateEndpoint)
		aggregateEndpoint = InstrumentingMiddleware(duration.With("method", "Aggregate"))(aggregateEndpoint)
	}

	var calculateEndpoint = MakeCalculateEndpoint(svc)
	{
		calculateEndpoint = ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Every(time.Second), 1))(calculateEndpoint)
		calculateEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(calculateEndpoint)
		calculateEndpoint = LoggingMiddleware(log.With(logger, "method", "Calculate"))(calculateEndpoint)
		calculateEndpoint = InstrumentingMiddleware(duration.With("method", "Calculate"))(calculateEndpoint)
	}

	return Set{
		AggregateEndpoint: aggregateEndpoint,
		CalculateEndpoint: calculateEndpoint,
	}
}

type AggregateRequest struct {
	Value float64 `json:"value"`
	OBUID int     `json:"obuId"`
	Unix  int64   `json:"unix"`
}

type CalculateRequest struct {
	OBUID int `json:"obuId"`
}

type CalculateResponse struct {
	OBUID         int     `json:"obuId"`
	TotalDistance float64 `json:"totalDistance"`
	TotalAmount   float64 `json:"totalAmount"`
	Err           error   `json:"error"`
}

type AggregateResponse struct {
	Err error `json:"error"`
}

func (s Set) Aggregate(ctx context.Context, distance types.Distance) error {
	_, err := s.AggregateEndpoint(ctx, AggregateRequest{
		OBUID: distance.OBUID,
		Value: distance.Value,
		Unix:  distance.Unix,
	})

	return err
}

func (s Set) Calculate(ctx context.Context, obuId int) (*types.Invoice, error) {
	resp, err := s.CalculateEndpoint(ctx, CalculateRequest{
		OBUID: obuId,
	})

	if err != nil {
		return nil, err
	}
	result := resp.(CalculateResponse)
	return &types.Invoice{
		OBUID:         result.OBUID,
		TotalDistance: result.TotalDistance,
		InvoiceAmount:   result.TotalAmount,
	}, nil
}

func MakeAggregateEndpoint(s aggservice.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(AggregateRequest)
		err = s.Aggregate(ctx, types.Distance{
			OBUID: req.OBUID,
			Value: req.Value,
			Unix:  req.Unix,
		})

		return AggregateResponse{Err: err}, err
	}
}

func MakeCalculateEndpoint(s aggservice.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(CalculateRequest)
		invoice, err := s.Calculate(ctx, req.OBUID)

		return CalculateResponse{
			OBUID:         invoice.OBUID,
			TotalDistance: invoice.TotalDistance,
			TotalAmount:   invoice.InvoiceAmount,
			Err:           err,
		}, err
	}
}