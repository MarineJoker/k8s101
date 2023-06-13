FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/marineJ/httpServer
COPY ./src $GOPATH/src/marineJ/httpServer
RUN go build .

EXPOSE 8080
ENTRYPOINT ["./httpServer"]
