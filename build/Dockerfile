FROM alpine:latest

ENV APP_PORT=80
ENV PUBLIC_PATH=/app/public
ENV LOG_PATH=/app/log
ENV MIN_MINUTES=60
ENV MAX_MINUTES=525600

COPY . /app
RUN mkdir -p /app/log

WORKDIR /app
ENTRYPOINT ["/app/http"]