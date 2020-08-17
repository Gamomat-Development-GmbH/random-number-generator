FROM golang:1.11.10-alpine3.9 as build

WORKDIR /workdir

RUN apk add --no-cache curl build-base git ca-certificates

COPY . /workdir

RUN echo $(go version) | echo $(ls -la)
RUN go build -o /app

FROM alpine:3.9
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /app /usr/local/bin/app
COPY scripts/ /usr/local/bin/

RUN echo 'checksum' & getAppHash.sh

ENTRYPOINT ["/usr/local/bin/app"]
