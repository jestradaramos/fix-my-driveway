FROM golang

ARG app_env
ENV APP_ENV app_env
ENV GO111MODULE=on

COPY . /go/src/github.com/jestradaramos/fix-my-driveway
WORKDIR /go/src/github.com/jestradaramos/fix-my-driveway

RUN apt-get update \
    && go get \
      golang.org/x/lint/golint 2>&1

RUN go mod vendor
RUN go mod tidy
