# 打包依赖阶段使用golang作为基础镜像
FROM golang:1.18 as builder

# 启用go module
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

# 指定OS等，并go build
RUN GOOS=linux GOARCH=amd64 go build ./cmd/node-runner

FROM ubuntu:22.04

WORKDIR /app
RUN apt-get update && apt-get -y install wget xvfb x11-apps graphicsmagick-imagemagick-compat&&\
     wget -q https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb \
    && apt-get install -y ./google-chrome-stable_current_amd64.deb

# 将上一个阶段publish文件夹下的所有文件复制进来
COPY --from=builder /app/node-runner .

EXPOSE 8080

ENTRYPOINT ["./node-runner"]
