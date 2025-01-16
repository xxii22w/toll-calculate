package main

import (
	"net"
	"net/http"
	"os"
	"tolling/go-kit-example/aggsrv/aggendpoint"
	"tolling/go-kit-example/aggsrv/aggservice"
	"tolling/go-kit-example/aggsrv/aggtransport"

	"github.com/go-kit/log"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)

	service := aggservice.New(logger)
	endpoints := aggendpoint.New(service, logger)
	httpHandler := aggtransport.NewHTTPHandler(endpoints, logger)

	// The HTTP listener mounts the Go kit HTTP handler we created.
	httpListener, err := net.Listen("tcp", ":4000")
	if err != nil {
		logger.Log("transport", "HTTP", "during", "Listen", "err", err)
		os.Exit(1)
	}
	logger.Log("transport", "HTTP", "addr", ":4000")
	err = http.Serve(httpListener, httpHandler)
	if err != nil {
		panic(err)
	}
}
