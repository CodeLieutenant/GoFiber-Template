ARG BASE_IMAGE="golang:1.24"
ARG PRODUCTION_IMAGE="debian:bookworm-slim"
ARG APP_NAME="GoFiber-Boilerplate"

FROM $PRODUCTION_IMAGE AS base-production

ENV DEBIAN_FRONTEND="noninteractive"
ENV GODEBUG="default=go1.24,cgocheck=0,disablethp=0,panicnil=0,http2client=1,http2server=1,asynctimerchan=0,madvdontneed=0"
ENV PATH="/usr/local/bin:${PATH}"

RUN apt-get update && apt-get upgrade -y \
	&& apt-get install -y ca-certificates \
    && apt-get autoremove -y \
    && apt-get clean \
	&& rm -rf /var/lib/apt/lists/*

EXPOSE 5000
EXPOSE 2112
EXPOSE 8080

ENTRYPOINT [ "/bin/bash" ]
CMD [ "${APP_NAME}", "serve" ]

FROM base-production AS production-goreleaser

COPY ${APP_NAME} /usr/local/bin/${APP_NAME}

FROM base-production AS production

COPY --from=build /app/bin/${APP_NAME} /usr/local/bin/${APP_NAME}
COPY --from=build /app/bin/config.yml /etc/${APP_NAME}/config.yml


FROM $BASE_IMAGE AS base

ENV GO111MODULE=on
ENV CFLAGS="-O0"
ENV CXXFLAGS="-O0"
ENV DEBIAN_FRONTEND="noninteractive"
ENV TZ="UTC"
ENV PATH="/usr/local/bin:${PATH}"

RUN apt-get update \
    && apt-get upgrade -y \
    && apt-get install iputils-ping mlocate vim ca-certificates make curl libc-dev -y \
    && curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b /usr/local/bin \
    && sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d -b /usr/local/bin \
    && updatedb

WORKDIR /app

FROM base AS build

ENV GOAMD64=v3
ENV GOARM64=v8.3,crypto
ENV CFLAGS="-O3"
ENV CXXFLAGS="-O3"

ARG VERSION

COPY . .

RUN task build VERSION=${VERSION}

FROM base AS develop

COPY . .

EXPOSE 2345
EXPOSE 80
EXPOSE 5000
EXPOSE 3000

ENTRYPOINT ["/bin/bash"]
CMD ["air", "-c", ".air.toml"]
