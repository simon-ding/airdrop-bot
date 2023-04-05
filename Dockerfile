# 打包依赖阶段使用golang作为基础镜像
FROM golang:1.20 as builder

# 启用go module
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

# 指定OS等，并go build
RUN GOOS=linux GOARCH=amd64 go build ./cmd/ -o airdrop-bot

FROM ubuntu:22.04

WORKDIR /app

# 将上一个阶段publish文件夹下的所有文件复制进来
COPY --from=builder /app/airdrop-bot .

EXPOSE 8080

ENTRYPOINT ["./airdrop-bot"]
