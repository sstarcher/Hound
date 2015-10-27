FROM gliderlabs/alpine

RUN apk-install ca-certificates git openssh

ADD hound /go/bin/houndd
ONBUILD COPY config.json /hound/

EXPOSE 6080

ENTRYPOINT ["/go/bin/houndd", "-conf", "/hound/config.json"]
