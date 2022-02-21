bin = "http"

# Make unit test
test:
	@go test ./cmd/$(bin)

# Build binary
build:
	@GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" -o ./dist/${bin} -v ./cmd/${bin}

# Run server
run:
	@make build
	@./dist/${bin}

# Clear dist
clear:
	@rm -f ./dist/${bin}