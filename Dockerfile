FROM golang:1.12.1

WORKDIR /app/src
ADD . .

ENV GO111MODULE=on

RUN go get github.com/pilu/fresh
CMD ["fresh"]
