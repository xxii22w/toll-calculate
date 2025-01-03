gate:
	@go build -o bin/gate gateway/main.go
	@./bin/gate


obu:
	@go build -o bin/obu obu/main.go
	@./bin/obu

receive:
	@go build -o bin/receive ./data_receiver
	@./bin/receive

calculator:
	@go build -o bin/calculator ./distance_calculator
	@./bin/calculator

agg:
	@go build -o bin/agg ./aggregator
	@./bin/agg

proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative  types/ptypes.proto

.PHONY: obu invoicer