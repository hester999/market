
FROM golang:1.24.3 AS builder

WORKDIR /app


RUN apt-get update && apt-get install -y make


COPY go.mod go.sum ./
RUN go mod download

# Копируем весь код
COPY . .


RUN go install github.com/swaggo/swag/cmd/swag@latest


RUN make swag && make build


FROM ubuntu:22.04

WORKDIR /app


RUN apt update && apt install -y ca-certificates


COPY --from=builder /app/cmd/app /app/app
COPY --from=builder /app/docs /app/docs
COPY --from=builder /app/static /app/static

EXPOSE 8080

CMD ["/app/app"]
