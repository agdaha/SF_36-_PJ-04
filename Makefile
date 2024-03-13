all: web-build build

build: 
	go build -o ./cmd/gonewServer -v ./cmd/main.go  

test:
	go test -v -race -timeout 30s ./...

web-modules-install:
	npm --prefix webapp install

web-build:
	npm --prefix webapp run build     
	rm -rf cmd/webapp
	cp -r webapp/dist cmd/webapp      

.DEFAULT_GOAL := build