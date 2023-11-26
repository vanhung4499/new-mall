FROM golang:1.21 as builder

WORKDIR /app
COPY . .
RUN go mod tidy
WORKDIR /app/cmd
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o ../main
WORKDIR /app

ENTRYPOINT ["./main"]