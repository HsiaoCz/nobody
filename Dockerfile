FROM golang:1.21.4 AS builder

COPY . /go/src

WORKDIR /go/src

