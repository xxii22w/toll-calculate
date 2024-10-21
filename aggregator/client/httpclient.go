package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"tolling/types"

	"github.com/sirupsen/logrus"
)

type HTTPClient struct {
	Endpoint string
}

func NewHTTPClient(endpoint string) *HTTPClient {
	return &HTTPClient{
		Endpoint: endpoint,
	}
}

func (c *HTTPClient) GetInvoice(ctx context.Context, invreq *types.GetInvoiceRequest) (*types.Invoice, error) {
	var (
		request = c.Endpoint + "/invoice?obu=" + strconv.Itoa(int(invreq.ObuID))
	)

	logrus.Infof("request -> %s", request)

	req, err := http.NewRequest("GET", request, nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("the service responsed with a status code of %d", resp.StatusCode)
	}

	var inv types.Invoice
	if err = json.NewDecoder(resp.Body).Decode(&inv); err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return &inv, nil
}

func (c *HTTPClient) AggregateDistance(ctx context.Context, aggReq *types.AggregatorDistanceRequest) error {
	b, err := json.Marshal(aggReq)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", c.Endpoint+"/aggregate", bytes.NewReader(b))
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("the service responded with non 200 status code %d", resp.StatusCode)
	}
	defer resp.Body.Close()
	return nil
}
