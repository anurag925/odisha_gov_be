# syntax=docker/dockerfile:1

FROM golang:1.19-alpine

RUN apk add make
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .
RUN go build -o myapp

# This container exposes port 3000 to the outside world
EXPOSE 3000

CMD ["make", "run"]