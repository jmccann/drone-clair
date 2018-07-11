# Docker image for the Drone clair plugin
#
#     cd $GOPATH/src/github.com/jmccann/drone-clair
#     go build
#     docker build --rm=true -t jmccann/drone-clair .

FROM alpine:3.8

RUN apk update && \
  apk add \
    ca-certificates && \
  rm -rf /var/cache/apk/*

ADD https://github.com/optiopay/klar/releases/download/v2.2.0/klar-2.2.0-linux-amd64 /usr/local/bin/klar
RUN chmod 0755 /usr/local/bin/klar

ADD drone-clair /bin/
ENTRYPOINT ["/bin/drone-clair"]
