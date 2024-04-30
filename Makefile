## build: build golang
.PHONY: build
build:
	go build -o bin/server cmd/server/server.go

## start: run server in development mode
.PHONY: start
start:
	./scripts/entrypoint.dev.sh

.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo