BINARY_NAME = tweet_microservice
BUILD_DIR = build

.PHONY: build run clean

build:
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd/app

run: build
	@./$(BUILD_DIR)/$(BINARY_NAME)

clean:
	@rm -rf $(BUILD_DIR)

up:
	docker compose --env-file .env \
		-f internal/deploy/docker-compose.yaml \
		up --build
