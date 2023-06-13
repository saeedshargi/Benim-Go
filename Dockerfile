FROM golang:1.20

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go mod download && go mod verify

RUN go build -o main cmd/main.go

CMD ["/app/main"]