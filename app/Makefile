APP_NAME=app
OUTPUT_PATH=build

.PHONY: build run clean swag

build:
	mkdir -p $(OUTPUT_PATH)
	go build -o $(OUTPUT_PATH)/$(APP_NAME) ./cmd/main.go

run: build
	./$(OUTPUT_PATH)/$(APP_NAME)

clean:
	rm -rf $(OUTPUT_PATH)

swag:
	swag init --generalInfo cmd/main.go --parseInternal --output ./docs
