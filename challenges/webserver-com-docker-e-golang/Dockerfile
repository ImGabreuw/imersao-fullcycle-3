FROM golang:1.16

WORKDIR /go/src

COPY go.mod .

RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 8000

CMD ["./main"]