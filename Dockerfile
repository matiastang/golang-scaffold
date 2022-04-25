FROM golang:alpine
LABEL version="1.0.0"
LABEL maintainer="matias"
LABEL description="项目描述"
# docker中的工作目录
WORKDIR $GOPATH/src/go_gin_docker
# 将当前目录同步到docker工作目录下，也可以只配置需要的目录和文件（配置目录、编译后的程序等）
ADD ./src ./
# 由于所周知的原因，某些包会出现下载超时。这里在docker里也使用go module的代理服务
ENV GO111MODULE=on
ENV GOPROXY="https://goproxy.io"
# 指定编译完成后的文件名，可以不设置使用默认的，最后一步要执行该文件名
RUN go build -o go_gin_docker .
EXPOSE 8080
# 这里跟编译完的文件名一致
ENTRYPOINT  ["./go_gin_docker"]
# docker run --name test_go_gin_docker -p 8080:8080 -d go_gin_docker
RUN echo 'go_gin_docker images build ok!'