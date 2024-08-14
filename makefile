
run: build
	./.bin/app 
build: clear
	@go build -o ./.bin/app ./cmd/main.go
clear:
	@rm -rf ./bin/*