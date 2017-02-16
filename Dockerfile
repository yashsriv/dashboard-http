FROM golang:latest

MAINTAINER Yash Srivastav

RUN curl https://glide.sh/get | sh

ENV GOPATH /go
RUN mkdir -p $GOPATH/src
RUN mkdir -p $GOPATH/bin

ENV SRCPATH /go/src/github.com/yashsriv/dashboard-http
RUN mkdir -p $SRCPATH
WORKDIR $SRCPATH

RUN bash -c "git config --global http.followRedirects true"
RUN bash -c "echo 192.30.253.113 github.com >> /etc/hosts"

COPY glide.yaml $SRCPATH
COPY glide.lock $SRCPATH
COPY loop.sh $SRCPATH
RUN ./loop.sh


COPY . $SRCPATH
RUN cd $SRCPATH && go install

EXPOSE 8080

ENTRYPOINT ["/go/bin/dashboard-http"]
