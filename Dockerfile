FROM golang:1.26.0-alpine AS builder

ENV PORT=6767
ENV DATABASE_URL="./db/database.db"

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o main main.go

FROM scratch

COPY --from=builder /app/main /main
COPY --from=builder /app/frontend ./frontend
COPY --from=builder /app/db ./db

EXPOSE 6767

CMD ["/main"]
