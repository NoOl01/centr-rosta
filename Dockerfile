FROM golang:1.25-alpine AS backend

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o /backend ./cmd/app

FROM alpine:latest

WORKDIR /app

COPY --from=backend /backend /app/backend

RUN chmod +x /app/backend

CMD ["/app/backend"]