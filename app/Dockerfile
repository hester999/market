FROM golang:1.24.3 AS builder

WORKDIR /app

RUN apt-get update && apt-get install -y make

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go install github.com/swaggo/swag/cmd/swag@latest

RUN make swag && make build


FROM ubuntu:22.04

WORKDIR /app/build

RUN apt update && apt install -y ca-certificates

# Копируем бинарник и статику
COPY --from=builder /app/build/app /app/build/app
COPY --from=builder /app/docs /app/docs
COPY --from=builder /app/static /app/static

EXPOSE 8080

CMD ["/app/build/app"]
