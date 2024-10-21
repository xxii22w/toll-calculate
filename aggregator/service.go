package main

import (
	"math/rand/v2"
	"tolling/types"

	"github.com/sirupsen/logrus"
)

const basePrice = 3.15

type Aggregator interface {
	AggregateDistance(types.Distance) error
	Invoice(int) (*types.Invoice, error)
}

type InvoiceAggregator struct {
	store Storer
}

func NewAggregator(store Storer) Aggregator {
	return &InvoiceAggregator{
		store: store,
	}
}

func (i *InvoiceAggregator) AggregateDistance(distance types.Distance) error {
	logrus.WithFields(logrus.Fields{
		"obuid":    distance.OBUID,
		"distance": distance.Value,
		"unix":     distance.Unix,
	}).Info("aggregating distance")
	return i.store.Insert(distance)
}

func (i *InvoiceAggregator) Invoice(obuID int) (*types.Invoice, error) {
	distSum, err := i.store.GetDistanceSum(obuID)
	if err != nil {
		return nil, err
	}
	return types.NewInvoice(obuID, distSum, rand.Float64()*distSum), nil
}
