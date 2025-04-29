dev:
	@go run cmd/server/main.go

dev-watch:
	@nodemon --watch . --ext go --exec go run cmd/server/main.go --signal SIGTERM