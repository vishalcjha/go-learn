IMAGE_NAME = playground
CONTAINER_NAME = playground
PORTS = -p 9998:9998
VOLUMES = -v ${PWD}:/go/src/github.com/MediaMath/playground
ENV GOPATH=/go/src/github.com/MediaMath/playground

help:
	@echo "Please use 'make <target>' where <target> is one of"
	@echo "  dbuild           build the docker image"
	@echo "  drebuild         rebuilds the image from scratch without using any cached layers"
	@echo "  drun             run the built docker image"
	@echo "  drestart         restarts the docker image"
	@echo "  dbash            starts bash inside a running container."
	@echo "  dtest            runs tests."

dbuild:
	@echo "Building docker image..."
	docker build -t ${IMAGE_NAME} .

drebuild:
	@echo "Rebuilding docker image..."
	docker build --no-cache=true -t ${IMAGE_NAME} .

drun:
	make dbuild
	docker run --rm=true --name=$(CONTAINER_NAME) $(VOLUMES) $(PORTS) $(LINK) -i -t ${IMAGE_NAME}

drestart: drun

dbash:
	docker exec -it $(CONTAINER_NAME) /bin/bash

dtest:
	make dbuild
	docker run --rm=true --name=$(CONTAINER_NAME) $(VOLUMES) $(PORTS) --entrypoint=/go/src/github.com/MediaMath/byoa-logbrain/test-entrypoint.sh ${IMAGE_NAME}

install:
	@echo "Installing Golang libraries ..."
	@go get -u github.com/kardianos/govendor
	@govendor init
	@govendor fetch +out