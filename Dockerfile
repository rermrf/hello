FROM golang:1.24

WORKDIR /build

COPY hello.go .

RUN go build -o hello hello.go

CMD ["./hello"]
