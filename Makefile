VERSION := $(shell git describe --tags --exact-match 2>/dev/null || echo latest)
DOCKERHUB_NAMESPACE ?= jormungandrk
IMAGE := ${DOCKERHUB_NAMESPACE}/microservice-user-profile:${VERSION}

build:
	docker build -t ${IMAGE} .

push: build
	docker push ${IMAGE}

run: build
	docker run ${ARGS} ${IMAGE}
