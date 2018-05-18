FROM golang:1.10
COPY . /blog_go
WORKDIR . /blog_go
RUN go get -u github.com/golang/dep/cmd/dep
RUN $GOPATH/bin/dep ensure
RUN go build src/

ENTRYPOINT ./go_build_src
EXPOSE 9345
