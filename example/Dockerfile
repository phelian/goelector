ARG PKG_VER_ALPINE=3.12

FROM alpine:$PKG_VER_ALPINE

RUN apk add --no-cache ca-certificates

ADD elector elector
ENTRYPOINT ["/elector"]