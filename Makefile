TAG=eligundry/go-poker
VOLUME="$(PWD)":/usr/src/go-poker
DOCKER_RUN=docker run --rm -it -v $(VOLUME) -w /usr/src/go-poker $(TAG)

build:
	docker build -t $(TAG) .

pull:
	docker pull $(TAG)

test:
	$(DOCKER_RUN) go test

run:
	$(DOCKER_RUN) go poker.go
