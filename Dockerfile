FROM brossquad/fiber-dev:1.0.3 as debug
ARG PORT=4000
COPY . /app
WORKDIR /app
EXPOSE ${PORT} 40000
CMD ["/usr/bin/start-server", "/app/cli/server", "/app"]


FROM brossquad/fiber-dev:1.0.3 as dev
COPY . /app
WORKDIR /app
RUN taks build
EXPOSE 4000
CMD ["task", "dev" ]


FROM brossquad/fiber-dev:1.0.3 as builder
COPY . /app
WORKDIR /app
RUN task build-prod


FROM alpine:3.12 as production
ARG PORT=4000
COPY --from=builder /app/build /app
WORKDIR /app
EXPOSE ${PORT}
CMD ["./app"]

