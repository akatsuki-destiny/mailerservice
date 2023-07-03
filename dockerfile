FROM golang:1.20.5-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o mailerservice ./cmd/main.go

RUN chmod +x /app/mailerservice

FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/mailerservice /app
COPY ./dev.env ./dev.env

EXPOSE 3002:3002

CMD ["/app/mailerservice"]