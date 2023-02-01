FROM golang:1.19-alpine as builder

RUN apk update && apk add make git sudo mc
RUN mkdir -p /build
COPY ./ /build/
WORKDIR /build/
RUN go mod tidy
RUN go build -o ./urlapi ./cmd/api

ENTRYPOINT ["./urlapi"]

EXPOSE 443
