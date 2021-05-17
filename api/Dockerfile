FROM golang:1.16 AS builder

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o cryptoPair .

FROM alpine:3.10

WORKDIR /root/

RUN apk --no-cache add ca-certificates

RUN mkdir Config

COPY --from=builder /app/cryptoPair .
COPY --from=builder /app/Config /root/Config

CMD ["./cryptoPair"]