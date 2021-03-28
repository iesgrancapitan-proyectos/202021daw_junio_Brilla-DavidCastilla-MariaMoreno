FROM golang:1.16.1 as builder

WORKDIR /var/src

COPY ./go.mod .
COPY ./go.sum .

RUN go mod download

COPY . .

RUN go build ./cmd/server/main.go -o brilla

FROM alpine

COPY --from=builder /var/src/brilla /brilla

CMD ["./brilla"]

