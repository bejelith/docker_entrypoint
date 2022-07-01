FROM golang:1.18.3-alpine as source

WORKDIR /tmp/build

COPY ./ ./

RUN apk add --upgrade make gcc && make

FROM alpine

COPY --from=source /tmp/build/bin/docker_entrypoint /docker_entrypoint

ENTRYPOINT ["/docker_entrypoint"]
