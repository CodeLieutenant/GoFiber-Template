FROM golang:1.15-alpine as builder

COPY . /app

WORKDIR /app

RUN curl -sL https://taskfile.dev/install.sh | sh && task build-prod

FROM alpine:3.13

COPY --from=builder /app/build /app

WORKDIR /app

CMD ['./app']
