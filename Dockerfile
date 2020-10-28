FROM brossquad/fiber-dev:1.0.0 as builder
COPY . /app
WORKDIR /app
RUN task build-prod

FROM alpine:3.13 as dev
COPY --from=builder /app/build /app/build
COPY . /app
WORKDIR /app
CMD ['task', 'dev' ]

FROM brossquad/fiber-dev:1.0.0 as production
COPY --from=builder /app/build /app
WORKDIR /app
CMD ['./app']
