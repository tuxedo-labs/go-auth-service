SERVICE_NAME=auth-service
DOCKER_IMAGE=yourdockerhubusername/$(SERVICE_NAME)

build:
	go build -o $(SERVICE_NAME) cmd/main.go

docker-build:
	docker build -t $(DOCKER_IMAGE):latest .

test:
	go test ./...

docker-push:
	docker push $(DOCKER_IMAGE):latest

clean:
	rm -f $(SERVICE_NAME)
