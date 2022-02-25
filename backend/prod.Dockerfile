# Build
#######
FROM golang:1.16.1 as builder

WORKDIR /var/src

COPY ./go.mod .
COPY ./go.sum .

RUN go mod download

COPY . .

RUN go build -tags prod -o brilla ./cmd/server/main.go

# Run
#####
FROM debian:stable-slim

COPY --from=builder /var/src/brilla /brilla

CMD ["./brilla"]

