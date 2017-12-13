FROM golang:alpine

RUN apk update \
    && apk upgrade \
    && apk add --no-cache \
        bash \
        git \
        openssh

RUN go-wrapper download \
        github.com/adamclerk/deck\
        github.com/stretchr/testify/assert
RUN go-wrapper install \
        github.com/adamclerk/deck \
        github.com/stretchr/testify/assert

ADD . /usr/src/go-poker
WORKDIR /usr/src/go-poker
