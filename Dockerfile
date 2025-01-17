# FROM golang:1.23.4-alpine AS build

# Support CGO and SSL
# RUN apk --no-cache add gcc g++ make
# RUN apk add git

# ENV GO111MODULE=on \
#     CGO_ENABLED=0 \
#     GOOS=linux \
#     GOARCH=amd64

# WORKDIR /go/src/app

# COPY go.mod go.sum .env ./

# Menginstal semua package yang ada di go.mod
# Install dependencies with go mod
# RUN go mod tidy

# COPY . .  

# Build the app
# RUN GOOS=linux go build -ldflags="-s -w" -o ./bin/test ./main.go

# FROM alpine:3.21.2
# RUN apk --no-cache add ca-certificates
# WORKDIR /usr/bin
# COPY --from=build /go/src/app/bin /go/bin
# EXPOSE 8080
# ENTRYPOINT /go/bin/test --port 8080



FROM golang:1.23.4-alpine

# Menambahkan skrip wait-for-it.sh
RUN apk add --no-cache curl
RUN curl -o /usr/local/bin/wait-for-it.sh https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh && \
    chmod +x /usr/local/bin/wait-for-it.sh

# Menambahkan file aplikasi
ADD . /app
WORKDIR /app

# Salin go.mod dan go.sum terlebih dahulu untuk optimalkan caching layer
COPY go.mod go.sum ./

RUN go mod tidy

# Salin seluruh file proyek
COPY . .

# Build aplikasi Go
RUN go build -o main .

# Ekspose port aplikasi
EXPOSE 8080

# Perintah untuk menunggu hingga database siap dan kemudian menjalankan aplikasi
CMD ["/app/main"]

