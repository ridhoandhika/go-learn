
FROM golang:1.23.4-alpine AS builder

RUN apk --no-cache add gcc g++ make
RUN apk add git

WORKDIR /app
ADD . /app

ENV SERVER_HOST=goservice
ENV SERVER_PORT=8080
ENV DATABASE_HOST=pgsql
ENV DATABASE_PORT=5432
ENV DATABASE_USER=postgres
ENV DATABASE_PASSWORD=postgres
ENV DATABASE_NAME=postgres

RUN cd /app & go mod download
RUN cd /app & CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o goservice main.go

FROM alpine
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
WORKDIR /app
COPY --from=builder /app/goservice /app
COPY --from=builder /app /app

EXPOSE 8080

ENTRYPOINT ./goservice
