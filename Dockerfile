FROM golang:1.17-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
RUN go build -o /app/app

FROM alpine
WORKDIR /
COPY --from=builder /app/app /app
EXPOSE 4000
ENTRYPOINT ["/app"]
