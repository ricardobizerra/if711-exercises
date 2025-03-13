FROM golang:1.23

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN mkdir -p /app/data && chmod -R 777 /app/data

RUN go build -o app main.go