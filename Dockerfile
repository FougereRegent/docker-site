FROM golang:1.22.1 AS builder

WORKDIR /build

COPY . .

RUN go build -o bin main.go

FROM golang:1.22.1 AS app

WORKDIR /app

ENV GIN_MODE="release"
ENV DOCKER_SITE_SESSION_TIME=""
ENV DOCKER_SITE_DATABASE_TYPE="LOCAL"
ENV DOCKER_SITE_CONNECTION_STRING=""

COPY ./assets/ ./assets/
RUN true
COPY ./templates/ ./templates/
RUN true
COPY --from=builder /build/bin .

EXPOSE 8080
ENTRYPOINT ["./bin"]
