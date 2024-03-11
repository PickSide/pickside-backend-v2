# syntax=docker/dockerfile:1

FROM golang:1.21.6

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o ./bin/pickside-service cmd/api/main.go

EXPOSE 8080

CMD [ "./bin/pickside-service" ]