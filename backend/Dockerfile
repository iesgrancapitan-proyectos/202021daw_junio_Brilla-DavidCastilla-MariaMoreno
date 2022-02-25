FROM golang:1.16.1

RUN go get -u github.com/cosmtrek/air

WORKDIR /var/src

COPY go.mod .
COPY go.sum .

RUN go mod download

CMD ["air", "-c", ".air.toml"]

