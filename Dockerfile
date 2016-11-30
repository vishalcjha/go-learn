FROM golang:1.7

RUN apt-get update
RUN apt-get install -y ca-certificates

RUN apt-get install --fix-missing -y \
apache2-utils

# Get govendor go package manager
RUN go get -u github.com/kardianos/govendor

# Expose port
EXPOSE 9998

WORKDIR /go/src/github.com/MediaMath/playground
ENV GOPATH="/go"

#Define comand
ENTRYPOINT govendor init && govendor fetch +out && bash