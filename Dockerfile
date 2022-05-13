FROM golang:1.17.8-alpine as builder

WORKDIR /work

ENV GO111MODULE="on"
ENV GOPROXY="https://proxy.golang.com.cn,direct"
ENV GOSUMDB=gosum.io+ce6e7565+AY5qEHUk/qmHc5btzW45JVoENfazw8LielDsaI+lEbq6
ENV CGO_ENABLED=0

# 缓存 mod 检索-那些不常更改的模块
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -tags=jsoniter -ldflags "-s -w" -o kube-box .


# FROM alpine:3.15.0 as runner
# FROM busybox:1.35.0 as runner
# FROM vukomir/busybox as runner
# FROM tianon/toybox as runner
# FROM busybox:latest as runner
# FROM sgfoot/busybox:v0.0.1 as runner
FROM sgfoot/busybox:v0.1.1 as runner

WORKDIR /work

RUN mkdir bin data logs

WORKDIR /work/bin

ENV DATA_PATH="/work/data/"

COPY --from=builder /work/kube-box /work/bin/
COPY --from=builder /work/data/ip.data /work/data/

ENTRYPOINT ["/work/bin/kube-box"]