FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod ./

COPY . .

RUN go build -o student-app .

FROM alpine:latest

RUN adduser -D -u 1000 appuser

WORKDIR /app

COPY --from=builder --chown=appuser:appuser /app/student-app .

USER appuser

EXPOSE 8080

CMD ["./student-app"]