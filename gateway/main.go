package main

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"strconv"
	"time"
	"tolling/aggregator/client"
	"tolling/types"

	"github.com/sirupsen/logrus"
)

type apiFunc func(w http.ResponseWriter, r *http.Request) error

func main() {
	listenAddr := flag.String("listenAddr", ":6000", "the listen address of the HTTP server")
	flag.Parse()
	var (
		client     = client.NewHTTPClient("http://localhost:4000") // endpoint of the aggregator service
		invHandler = newInvoiceHandler(client)
	)
	http.HandleFunc("/invoice", makeAPIFunc(invHandler.handleGetInvoice))
	logrus.Infof("gateway HTTP server running on port %s", *listenAddr)
	log.Fatal(http.ListenAndServe(*listenAddr, nil))
}

type InvoiceHandler struct {
	client client.Client
}

func newInvoiceHandler(c client.Client) *InvoiceHandler {
	return &InvoiceHandler{client: c}
}

func (h *InvoiceHandler) handleGetInvoice(w http.ResponseWriter, r *http.Request) error {
	obuID, err := strconv.Atoi((r.URL.Query().Get("obu")))
	if err != nil {
		return err
	}
	inv, err := h.client.GetInvoice(context.Background(), &types.GetInvoiceRequest{
		ObuID: int64(obuID),
	}) 
	if err != nil {
		return err
	}
	return writeJSON(w, http.StatusOK, inv)
}

func writeJSON(w http.ResponseWriter, code int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(v)
}

func makeAPIFunc(fn apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func(start time.Time) {
			logrus.WithFields(logrus.Fields{
				"took": time.Since(start),
				"uri":  r.RequestURI,
			}).Info("REQ :: ")
		}(time.Now())
		if err := fn(w, r); err != nil {
			writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
	}
}
