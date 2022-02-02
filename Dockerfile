FROM golang:1.15

ADD . /go/src/aav_logger
RUN go get github.com/revel/revel \
    && go get github.com/revel/cmd/revel \
    && go get aav_logger/...

ENTRYPOINT revel run aav_logger
WORKDIR /go/src
EXPOSE 9000
