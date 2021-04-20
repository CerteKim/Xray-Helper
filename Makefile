all: build

clean:
	@rm -rf ./bin/helper-*

build:
	CGO_ENABLED=0 GOOS=linux go build -o ./bin/helper-$(shell uname -m) -trimpath -ldflags "-s -w -buildid=" ./main
	# CGO_ENABLED=0 GOOS=linux GOARCH=mipsle go build -o ./bin/helper-mipsle -trimpath -ldflags "-s -w -buildid=" ./main
	# CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=7 go build -o ./bin/helper-armv7 -trimpath -ldflags "-s -w -buildid=" ./main

run: 
	@go run ./main
