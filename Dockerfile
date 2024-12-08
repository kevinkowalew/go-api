FROM golang:1.19-buster AS builder

WORKDIR /src

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o /api main.go
FROM alpine:3.15.0
COPY --from=builder /api /api
CMD ["/api"]
