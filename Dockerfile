FROM golang:1.12

WORKDIR /go/src/app

COPY myplugin.go .
COPY config.toml .

RUN git clone https://github.com/rwynn/monstache.git
RUN cd monstache && GO111MODULE=on go build && cp monstache ../monst

RUN GO111MODULE=on go build -buildmode=plugin -o myplugin.so myplugin.go

CMD ["./monst","-f","config.toml"]
