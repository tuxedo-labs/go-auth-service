SERVICE_NAME=auth_service

build:
	go build -o $(SERVICE_NAME) main.go

docker-build:
	docker build -t $(DOCKER_IMAGE):latest .

clean:
	rm -f $(SERVICE_NAME)

run:
    ./${SERVICE_NAME}
