FROM golang:alpine AS builder
LABEL stage=gobuilder
ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct
WORKDIR /app
ADD go.mod .
ADD go.sum .
RUN go mod download
COPY .. .
RUN go build -ldflags="-s -w" -o /app/server main.go
FROM alpine
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk update --no-cache
RUN apk add --no-cache ca-certificates
RUN apk add --no-cache tzdata
ENV TZ Asia/Shanghai
WORKDIR /app
COPY --from=builder /app/server /app/server
CMD ["./server"]