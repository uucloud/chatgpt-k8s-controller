# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=chatgpt-k8s-controller
BIN_DIR=bin
IMAGE_NAME=uucloud/chatgpt-k8s-contronller

.PHONY: all build clean test docker-build docker-push deploy

all: test build

gen:
	@echo "Clean generated..."
	rm -rf ./pkg/generated

	@echo "Generate code..."
	./hack/update-codegen.sh

build:
	@echo "Building binary..."
	CGO_ENABLED=0 $(GOBUILD) -o $(BIN_DIR)/$(BINARY_NAME) 

clean:
	@echo "Cleaning up..."
	$(GOCLEAN)
	rm -rf $(BIN_DIR)

test:
	@echo "Running tests..."
	$(GOTEST) -v ./...

docker-build:
	@echo "Building Docker image..."
	docker build -t $(IMAGE_NAME) .

docker-push:
	@echo "Pushing Docker image..."
	docker push $(IMAGE_NAME)

deploy:
	@echo "Deploying to Kubernetes..."
	kubectl apply -f artifacts/clusterrole.yaml
	kubectl apply -f artifacts/serviceaccount.yaml
	kubectl apply -f artifacts/clusterrolebinding.yaml
	kubectl apply -f artifacts/crd.yaml
	kubectl apply -f artifacts/manager.yaml

undeploy:
	@echo "Undeploying..."
	kubectl delete -f artifacts/manager.yaml
	kubectl delete -f artifacts/clusterrolebinding.yaml
	kubectl delete -f artifacts/clusterrole.yaml
	kubectl delete -f artifacts/serviceaccount.yaml
	kubectl delete -f artifacts/crd.yaml


