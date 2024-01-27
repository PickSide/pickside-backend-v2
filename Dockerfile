# syntax=docker/dockerfile:1

FROM golang:1.21.6

WORKDIR /me-service

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping

EXPOSE 8080

# Run
CMD ["/docker-gs-ping"]