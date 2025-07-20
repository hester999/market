APP_NAME=app

.PHONY: build run clean swag

build:
	go build -o ./cmd/$(APP_NAME) ./cmd/main.go

run: build
	./cmd/$(APP_NAME)

clean:
	rm -rf ./cmd/$(APP_NAME)

swag:
	swag init --generalInfo cmd/main.go --parseInternal --output ./docs
