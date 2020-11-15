.PHONY: build clean deploy

build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/createdevice createdevice/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/getdevice getdevice/main.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	sls deploy --verbose
