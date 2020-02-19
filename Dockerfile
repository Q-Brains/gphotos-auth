FROM golang:alpine

ENV LANG C.UTF-8
ENV WORK_DIR /usr/local/gphotos-auth

RUN apk add --no-cache git && \
    go get github.com/Q-Brains/gphotos && \
    apk del --purge git

WORKDIR ${WORK_DIR}

ADD . ${WORK_DIR}

CMD [ "go", "run", "main.go" ]
