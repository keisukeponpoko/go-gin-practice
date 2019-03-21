FROM golang:1.12.1

WORKDIR /app/src
ADD . .

ENV GO111MODULE=on

CMD ["go", "run", "main.go"]
