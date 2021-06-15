GNUMAKEFLAGS=--no-print-directory

start:
	@make build && \
	make run

build:
	@go build ./services/proxy/main.go

run:
	@./main
