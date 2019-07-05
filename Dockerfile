FROM golang:1.9

WORKDIR $GOPATH/src/xapi

#将服务器的go工程代码加入到docker容器中
ADD . $GOPATH/src/xapi

RUN go build .

EXPOSE 8000

ENTRYPOINT  ["./xapi"]
