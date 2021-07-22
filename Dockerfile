FROM golang:1.16 as mod
LABEL stage=mod
ARG GOPROXY=https://goproxy.cn,https://mirrors.aliyun.com/goproxy/,https://goproxy.io,direct
WORKDIR /root/app/

COPY go.mod ./
COPY go.sum ./
RUN go mod download


FROM mod as builder
LABEL stage=intermediate0
ARG LDFLAGS
ARG GOARCH=amd64
COPY ./ ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=${GOARCH} go build -o goWebDemo -ldflags "${LDFLAGS}" main.go

EXPOSE 3000

ENTRYPOINT ["/goWebDemo"]