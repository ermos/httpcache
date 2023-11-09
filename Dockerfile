FROM golang:1.21-alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY cmd ./cmd
COPY internal ./internal

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" -o http -v ./cmd/http

FROM alpine:latest

ARG APP_PORT=80
ARG PUBLIC_PATH=/app/public
ARG LOG_PATH=/app/log
ARG MIN_MINUTES=60
ARG MAX_MINUTES=525600

# index.html
ARG NAME=HTTPCache
ARG HOST=http://localhost:${APP_PORT}

ENV APP_PORT=$APP_PORT
ENV PUBLIC_PATH=$PUBLIC_PATH
ENV LOG_PATH=$LOG_PATH
ENV MIN_MINUTES=$MIN_MINUTES
ENV MAX_MINUTES=$MAX_MINUTES

WORKDIR /app

RUN mkdir -p $LOG_PATH $PUBLIC_PATH

COPY --from=builder /build/http /http
COPY public $PUBLIC_PATH

RUN sed -i "s/{{.Name}}/${NAME}/g" $PUBLIC_PATH/index.html
RUN sed -i "s@{{.Host}}@${HOST}@g" $PUBLIC_PATH/index.html

EXPOSE $APP_PORT

ENTRYPOINT ["/http"]