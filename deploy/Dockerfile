FROM golang:1.23.2-bookworm AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o goserver ./cmd/app/main.go

# Deploy
FROM alpine:3.20.3

WORKDIR /app

COPY --from=builder /app/goserver .

EXPOSE 8080

CMD ["./goserver"]