FROM golang:alpine

ENV LANG C.UTF-8
ENV WORK_DIR /usr/local/gphotos-auth

ENV CLIENT_ID=
ENV CLIENT_SECRET=
ENV AUTH_SCOPES=

RUN apk add --no-cache git && \
    go get github.com/Q-Brains/gphotos && \
    apk del --purge git

WORKDIR ${WORK_DIR}

ADD . ${WORK_DIR}
