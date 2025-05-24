FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o webapp

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/webapp .
COPY --from=builder /app/views ./views
COPY --from=builder /app/assets ./assets

EXPOSE 8080

CMD ["./webapp"]