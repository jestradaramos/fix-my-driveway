FROM golang

ARG app_env
ENV APP_ENV app_env
ENV GO111MODULE=on

COPY . /go/src/github.com/jestradaramos/fix-my-driveway
WORKDIR /go/src/github.com/jestradaramos/fix-my-driveway

RUN apt-get update \
  && go get \
  golang.org/x/tools/cmd/gopls \
  github.com/mdempsky/gocode \
  github.com/uudashr/gopkgs/cmd/gopkgs \
  github.com/ramya-rao-a/go-outline \
  github.com/acroca/go-symbols \
  golang.org/x/tools/cmd/guru \
  golang.org/x/tools/cmd/gorename \
  github.com/go-delve/delve/cmd/dlv \
  github.com/stamblerre/gocode \
  github.com/rogpeppe/godef \
  golang.org/x/tools/cmd/goimports \
  golang.org/x/lint/golint 2>&1 

RUN go mod vendor
RUN go mod tidy
RUN go build main.go
