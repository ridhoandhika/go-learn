FROM golang:1.23.4-alpine AS build

# Support CGO and SSL
RUN apk --no-cache add gcc g++ make
RUN apk add git

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /go/src/app

COPY go.mod go.sum .env ./

# Menginstal semua package yang ada di go.mod
# Install dependencies with go mod
RUN go mod tidy

RUN go mod download

COPY . .  

# Build the app
RUN GOOS=linux go build -ldflags="-s -w" -o ./bin/test ./main.go

FROM alpine:3.21.2
RUN apk --no-cache add ca-certificates
WORKDIR /usr/bin
COPY --from=build /go/src/app/bin /go/bin
EXPOSE 8080
ENTRYPOINT /go/bin/test --port 8080



