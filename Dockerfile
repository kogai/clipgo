FROM golang:1.8

ADD . $GOPATH/src/app
WORKDIR $GOPATH/src/app

RUN go get github.com/atotto/clipboard
RUN go get gopkg.in/alecthomas/kingpin.v2

CMD go run golip.go

