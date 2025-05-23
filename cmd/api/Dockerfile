FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ../../
RUN cd ../.. && go mod download

COPY . ../../

WORKDIR /app/cmd/api

RUN go build -o api

# Final image
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/cmd/api/api .

COPY ../../cmd/api/.env .env

EXPOSE 5000

CMD ["./api"]