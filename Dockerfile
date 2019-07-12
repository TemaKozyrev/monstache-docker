FROM golang:1.12 AS builder

WORKDIR /go/src/app

COPY myplugin.go .
COPY config.toml .

RUN git clone https://github.com/rwynn/monstache.git
RUN cp myplugin.go ./monstache/myplugin.go
RUN cd monstache && GO111MODULE=on go mod vendor
RUN cd monstache && GO111MODULE=on go build -mod=vendor -buildmode=plugin -o myplugin.so myplugin.go && cp myplugin.so ../myplugin.so

RUN cd monstache && GO111MODULE=on GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -v -mod=vendor && cp monstache ../monst


# RUN ls -la
# RUN monst
CMD ["./monst","-f","config.toml"]
# ENTRYPOINT ["sh","./monst","-f","config.toml"]
