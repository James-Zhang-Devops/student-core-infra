FROM golang:1.21-alpine

WORKDIR /app
COPY . .

RUN go build -o student-app .

CMD ["./student-app"]