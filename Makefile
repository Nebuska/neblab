# ------------------------------
# Variables
# ------------------------------
APP_NAME := TaskTracker
SRC := ./cmd/server/main.go
BIN := ./bin/${APP_NAME}
DOCKER_IMAGE := ${APP_NAME}:latest


# ------------------------------
# Default build and run
# ------------------------------
all: clean build

# ------------------------------
# Build the binary
# ------------------------------
build:
	@echo "Building ${APP_NAME}"
	go build -o ${BIN} ${SRC}

# ------------------------------
# Run the application
# ------------------------------
run: build
	@echo "Running ${APP_NAME}"
	$(BIN)

# ------------------------------
# Clean the bin
# ------------------------------
clean:
	@echo "Cleaning previous build"
	rm -rf ./bin/*

# ------------------------------
# Help
# ------------------------------
help:
	@echo "This is the help you need!"