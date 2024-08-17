build:
	go build -o bin/quantTradingEngine

run: build
	./bin/quantTradingEngine

test:
	go test -v ./...