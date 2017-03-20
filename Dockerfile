FROM golang:latest

MAINTAINER Yash Srivastav

RUN curl https://glide.sh/get | sh

ENV GOPATH /go
ENV GO_ENV production
RUN mkdir -p $GOPATH/src
RUN mkdir -p $GOPATH/bin
RUN go get github.com/markbates/pop/...
RUN go install github.com/markbates/pop/soda

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

ENTRYPOINT ["/go/src/github.com/yashsriv/dashboard-http/run.sh"]
