# Variables
BINARY_NAME=akaSocial
MAIN_FILE=main.go
BIN_DIR=bin
DB_MIGRATIONS_FOLDER=internal/db/migrations
DB_CONNECTION_STRING="postgres://user:password@localhost:5432/dbname?sslmode=disable"
#change if you want to test with your db :)

all: build

# Create the bin directory if it doesn't exist
$(BIN_DIR):
	mkdir -p $(BIN_DIR)

# Build the application
build: $(BIN_DIR)
	go build -o $(BIN_DIR)/$(BINARY_NAME) $(MAIN_FILE)

# Run the application
run: build
	./$(BIN_DIR)/$(BINARY_NAME)








clean:
	rm -f $(BINARY_NAME)

# Run database migrations
migrate-up:
	migrate -path $(DB_MIGRATIONS_FOLDER) -database $(DB_CONNECTION_STRING) up

migrate-down:
	migrate -path $(DB_MIGRATIONS_FOLDER) -database $(DB_CONNECTION_STRING) down


#docker
docker-build:
	docker build -t $(BINARY_NAME):latest .

docker-run:
	docker run -p 8080:8080 --env-file .env $(BINARY_NAME):latest

docker-clean:
	docker stop $(shell docker ps -aq) && docker rm $(shell docker ps -aq)
