FROM golang:1.19 as builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64\
    GOPROXY=https://goproxy.cn,direct 

WORKDIR /build
COPY . .
RUN go mod downloadd 
RUN go build  -ldflags="-w -s" -o main . -t mall1.0
RUN mkdir publish  \
    && cp main publish  \
    && cp -r conf publish

FROM busybox:1.28.4

WORKDIR /app

COPY --from=builder /app/publish .

# 指定运行时环境变量
ENV GIN_MODE=release
EXPOSE 3000

ENTRYPOINT ["./main"]