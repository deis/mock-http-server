FROM deis/go:latest

WORKDIR /go/src/github.com/deis/mock-http-server
ADD . /go/src/github.com/deis/mock-http-server
RUN CGO_ENABLED=0 go get -a -ldflags '-s' github.com/deis/mock-http-server
