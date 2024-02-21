FROM golang:alpine3.16 AS builder

RUN apk update && apk add --no-cache ca-certificates

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

RUN go build -o /OzonTest ./cmd/zip

FROM scratch

WORKDIR /bin

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder OzonTest OzonTest

ENTRYPOINT ["OzonTest"]
