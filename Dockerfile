FROM golang:1.18 as builder

WORKDIR /app
COPY . .
RUN go mod tidy
WORKDIR /app/cmd
RUN CGO_ENABLED=0 GOOS=linux go build  -ldflags="-w -s" -o ../main
WORKDIR /app
RUN mkdir publish  \
    && cp main publish  \
    && cp -r config publish

FROM busybox:1.28.4

WORKDIR /app

COPY --from=builder /app/publish .

#Specify runtime environment variables
ENV GIN_MODE=release
EXPOSE 5000

ENTRYPOINT ["./main"]