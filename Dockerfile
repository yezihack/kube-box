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

# 添加额外的工具包
RUN go install github.com/davecheney/httpstat@latest

FROM sgfoot/busybox:v0.1.1 as runner

WORKDIR /work

RUN mkdir bin data logs

WORKDIR /work/bin

ENV DATA_PATH="/work/data/"

COPY --from=builder /work/kube-box /work/bin/
COPY --from=builder /work/data/ip.data /work/data/
COPY --from=builder /go/bin/httpstat /usr/local/bin/

ENTRYPOINT ["/work/bin/kube-box"]