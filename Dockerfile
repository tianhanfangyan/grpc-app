FROM golang:1.12-alpine AS builder

ARG RELEASE
ARG COMMIT
ARG BUILD
ARG GOPROXY

ENV CGO_ENABLED=0
ENV GOPROXY=${GOPROXY}

WORKDIR /grpc-app

ADD . .

RUN export PROJECT=$(head -n1 go.mod | awk '{ print $2 }') && go build --mod=vendor\
        -ldflags "-s -w \
        -X ${PROJECT}/version.Release=${RELEASE} \
        -X ${PROJECT}/version.Commit=${COMMIT} \
        -X ${PROJECT}/version.Build=${BUILD}" \
        -o /tmp/grpc-app \
        .

FROM alpine:3.7

LABEL maintainer="dystargate@gmail.com"

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories && \
        apk add --no-cache -U tzdata ca-certificates

COPY --from=builder /tmp/grpc-app /grpc-app

ENTRYPOINT ["/grpc-app"]

