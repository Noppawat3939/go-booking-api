dev:
	@go run cmd/server/main.go

dev-watch:
	@echo "🚀🌐 running locally and watching in termial..."
	@nodemon --watch . --ext go --exec go run cmd/server/main.go --signal SIGTERM