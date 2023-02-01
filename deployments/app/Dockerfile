FROM golang:1.19-alpine as builder

RUN apk update && apk add make git sudo mc
RUN mkdir -p /build
COPY ./ /build/
WORKDIR /build/
RUN go mod tidy
RUN go build -o ./disp ./cmd/url-dispatcher
RUN go build -o ./api ./cmd/api

ENTRYPOINT ["/build/startprocess.sh"]

EXPOSE 443
