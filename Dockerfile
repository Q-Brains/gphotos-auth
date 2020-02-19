FROM golang:alpine

ENV LANG C.UTF-8
ENV WORK_DIR /usr/local/gphotos-auth

ENV CLIENT_ID=
ENV CLIENT_SECRET=
ENV SCOPES=

RUN apk add --no-cache && \
    go get github.com/Q-Brains/gphotos

VOLUME ${WORK_DIR}

WORKDIR ${WORK_DIR}
