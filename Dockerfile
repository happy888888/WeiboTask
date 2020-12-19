FROM golang:1.15.6-alpine3.12 as builder
WORKDIR /go
COPY ./src ./src
RUN CGO_ENABLED=0 GOPATH=$(pwd) go install -ldflags "-s -w" WeiboTask

FROM alpine:3.12
LABEL maintainer="星辰"

RUN apk update \
    && apk add tzdata \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone

COPY --from=builder /go/bin/WeiboTask /usr/bin/WeiboTask

ENTRYPOINT ["/usr/bin/WeiboTask", "-L"]
CMD ["-l", "/tmp/WeiboTask.log"]