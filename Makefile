build:
	go build -o ./bin/email ./cmd/email
	go build -o ./bin/server ./cmd/server
	go build -o ./bin/worker ./cmd/worker

.PHONY: build
