FROM golang:1.23-alpine AS builder

WORKDIR /app


COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go mod tidy

WORKDIR /app/cmd/api

RUN go mod tidy
RUN go build -o api

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/cmd/api/api .
EXPOSE 5000

CMD ["./api"]