
run: build
	./.bin/menv
build: clear
	@go build -o ./.bin/menv ./cmd/main.go
clear:
	@rm -rf ./bin/*
