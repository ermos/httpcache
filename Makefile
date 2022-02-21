bin = "http"

# Make unit test
test:
	@go test ./cmd/$(bin)

# Compile binary
compile:
	@GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" -o ./dist/${bin} -v ./cmd/${bin}

# Run server
run:
	@make compile
	@./dist/${bin}

# Clear dist
clear:
	@rm -f ./dist/${bin}