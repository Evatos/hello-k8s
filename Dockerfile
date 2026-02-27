FROM golang:1.26-alpine AS builder
WORKDIR /app
COPY go.mod ./
COPY main.go ./
RUN go build -o hello-server main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/hello-server .
EXPOSE 8080
CMD ["./hello-server"]