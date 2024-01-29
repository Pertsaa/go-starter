build:
	go build -o ./bin/cli ./cmd/cli
	go build -o ./bin/server ./cmd/server
	go build -o ./bin/worker ./cmd/worker

.PHONY: build
