all: get_deps build

get_deps:
	rm -rf ./build
	rm -rf ./vendor
	go mod download
	go mod vendor

build:
	go build -o build/exporter exporter.go
	go build -o build/cron cron-exporter.go
