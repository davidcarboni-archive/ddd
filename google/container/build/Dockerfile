FROM golang

WORKDIR /go/src
ADD main.go .

CMD go build -o google . && ls -lah && ./google
