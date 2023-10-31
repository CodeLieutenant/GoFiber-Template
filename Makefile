.PHONY: build
build:
ifeq ($(ENV),production)
	@CGO_ENABLED=0 CXX=g++ CC=gcc go build -ldflags="-s -w -X 'main.Version=${VERSION}'" -o ./bin/$(APP_NAME) ./main.go
else ifeq ($(ENV),development)
	@CXX=g++ CC=gcc go build -o ./bin/$(APP_NAME) -gcflags "all=-N -l" ./main.go
else
	@echo "Target ${ENV} is not supported"
endif
	@cp ./config.example.yml bin/config.yml
