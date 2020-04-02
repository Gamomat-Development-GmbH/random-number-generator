FROM golang:1.11.10-alpine3.9 as build

RUN apk add --no-cache curl build-base git ca-certificates
RUN go get -u github.com/seehuhn/mt19937

WORKDIR /go/src
COPY . code-fabrik.com/random-number-generator

RUN echo $(go version)
RUN go build -o /app code-fabrik.com/random-number-generator

FROM alpine:3.9
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /app /usr/local/bin/app
ENTRYPOINT ["/usr/local/bin/app"]
