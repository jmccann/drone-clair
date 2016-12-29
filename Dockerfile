# Docker image for the Drone clair plugin
#
#     cd $GOPATH/src/github.com/Unikorn123/drone-clair
#     go build
#     docker build --rm=true -t Unikorn123/drone-clair .

FROM alpine:3.3

RUN apk update && \
  apk add \
    ca-certificates && \
  rm -rf /var/cache/apk/*

ADD drone-clair /bin/
ENTRYPOINT ["/bin/drone-clair"]
