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

COPY ./wait-for-it.sh /wait-for-it.sh
RUN chmod +x /wait-for-it.sh

COPY --from=builder /var/src/brilla /brilla

CMD ["./wait-for-it.sh", "database:8529", "--", "./brilla"]

