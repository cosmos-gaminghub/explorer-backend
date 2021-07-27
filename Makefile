all: get_deps build

get_deps:
	rm -rf ./vendor
	go mod download
	go mod vendor

build:
	rm -rf ./build
	go build -o build/exporter exporter.go
	go build -o build/cron cron.go
