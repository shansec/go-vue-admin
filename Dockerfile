FROM golang:alpine as builder

WORKDIR /go/go-vue-admin
COPY . .

RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
    && go mod tidy \
    && go build -o server .

FROM alpine:latest

WORKDIR /go/go-vue-admin

COPY --from=0 /go/go-vue-admin ./

EXPOSE 8080

ENTRYPOINT ./server