FROM golang:1.17

WORKDIR /go/src

RUN apt-get update && apt-get install sqlite3 -y

RUN go install github.com/golang/mock/mockgen@v1.6.0

CMD ["tail", "-f", "/dev/null"]