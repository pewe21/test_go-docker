FROM golang:1.22-alpine

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o test_go

EXPOSE 8080

CMD ./test_go
