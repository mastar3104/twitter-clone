GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOFMT=$(GOCMD) fmt

start-docker:
	docker-compose -f ../db/docker-compose.yaml up -d
	sleep 60 # mySQLが立ち上がるまで待機
stop-docker:
	docker-compose -f ../db/docker-compose.yaml down
test:
	make start-docker
	$(GOTEST) -v ./...
	make stop-docker
build:
	$(GOBUILD)
clean:
	make stop-docker
	$(GOCLEAN)
run:
	$(GORUN) main.go
dev:
	make start-docker
	$(GORUN) main.go
fmt:
	$(GOFMT)