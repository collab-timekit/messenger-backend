# Definiowanie nazw dla plików i obrazów
BINARY_NAME=messenger
DOCKER_IMAGE=messenger:latest
DOCKERFILE=Dockerfile
PORT=8080

# Kompilacja aplikacji Go
go-build:
	@echo "Building Go application..."
	go build -o $(BINARY_NAME)

# Uruchomienie aplikacji Go lokalnie
go-run:
	@echo "Running Go application..."
	go run main.go

# Budowanie obrazu Docker
docker-build:
	@echo "Building Docker image..."
	docker build -t $(DOCKER_IMAGE) -f $(DOCKERFILE) .

# Uruchomienie aplikacji w Dockerze
docker-run:
	@echo "Running Docker container..."
	docker run -p $(PORT):$(PORT) $(DOCKER_IMAGE)

# Budowanie Go i uruchomienie aplikacji
all: go-build go-run

# Usunięcie skompilowanego pliku binarnego
clean:
	@echo "Cleaning up..."
	rm -f $(BINARY_NAME)