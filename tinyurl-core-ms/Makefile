BIN=./bin/microservice

.PHONY: build run test clean

build:
	@echo "Building the project..."
	go build -o $(BIN) ./cmd/microservice/main.go

run: build
	@echo "Running the microservice..."
	$(BIN)

test:
	@echo "Running tests..."
	go test ./tests...

clean:
	@echo "Cleaning up..."
	rm -f $(BIN)