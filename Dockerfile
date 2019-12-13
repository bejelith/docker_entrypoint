FROM golang:1.13.5-alpine3.10 as source

WORKDIR /tmp/build

COPY ./ ./

RUN apk add --upgrade make gcc && make

FROM alpine

COPY --from=source /tmp/build/bin/docker_entrypoint /docker_entrypoint

ENTRYPOINT ["/docker_entrypoint"]
