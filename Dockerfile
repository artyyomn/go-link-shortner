FROM golang:1.26.0-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download


COPY . .

RUN go build -o main .

FROM alpine:latest

COPY --from=builder /app/main /main

EXPOSE 6767

CMD["/main"]
