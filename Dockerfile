FROM golang:1.14

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /app
COPY . .
RUN go build .

EXPOSE 8080
ENTRYPOINT ["./toolsapi"]