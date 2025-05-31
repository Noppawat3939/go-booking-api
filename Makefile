dev:
	@go run cmd/api/main.go

dev-watch:
	@echo "🚀🌐 running locally and watching in termial..."
	@nodemon --watch . --ext go --exec go run cmd/api/main.go --signal SIGTERM