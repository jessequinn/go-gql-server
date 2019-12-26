FROM golang as base

WORKDIR /app

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

#COPY .dev.env .env
COPY . .

RUN go mod download

# it will take the flags from the environment
RUN go build -o build/gql-server cmd/gql-server/main.go

### Certs
FROM alpine:latest as certs
RUN apk --update add ca-certificates

### App
FROM scratch as app
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=base app /
ENTRYPOINT ["/build/gql-server"]

EXPOSE 7777
