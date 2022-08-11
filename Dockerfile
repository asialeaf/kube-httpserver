# Build the kube-httpserver binary
FROM golang:1.18 as builder

RUN apt-get -y update && apt-get -y install upx

WORKDIR /workspace
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum

# Copy the go source
COPY pkg/ pkg/
COPY cmd/ cmd/

# Build
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
ENV GO111MODULE=on
ENV GOPROXY="https://goproxy.cn,direct"

RUN go mod download && \
    go build -a -o kube-httpserver cmd/kube-httpserver/main.go && \
    upx kube-httpserver

FROM alpine:3.13
COPY --from=builder /workspace/kube-httpserver .
ENTRYPOINT ["/kube-httpserver"]
