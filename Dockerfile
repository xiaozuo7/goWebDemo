FROM golang:1.16.10 as mod
LABEL stage=mod
ARG GOPROXY=https://goproxy.cn,https://mirrors.aliyun.com/goproxy/,https://goproxy.io,direct
WORKDIR /root/goWebDemo

COPY go.mod ./
COPY go.sum ./
RUN go mod download

FROM mod as builder
LABEL stage=intermediate0
ARG LDFLAGS
ARG GOARCH=amd64
COPY ./ ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=${GOARCH} \
   go build -o goWebDemo \
   -gcflags="all=-trimpath=`pwd` -N -l" \
   -asmflags "all=-trimpath=`pwd`" \
   -ldflags "${LDFLAGS}" main.go


FROM alpine:3.13.5

COPY --from=builder /root/goWebDemo /goWebDemo
COPY --from=builder /root/goWebDemo/config /config

ENV TZ Asia/Shanghai

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
    apk update && \
    apk add --no-cache \
      curl \
      ca-certificates \
      bash \
      iproute2 \
      tzdata && \
    ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo Asia/Shanghai > /etc/timezone && \
    if [ ! -e /etc/nsswitch.conf ];then echo 'hosts: files dns myhostname' > /etc/nsswitch.conf; fi && \
   rm -rf /var/cache/apk/* /tmp/*

EXPOSE 3000
ENTRYPOINT ["/goWebDemo"]
