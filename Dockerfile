FROM golang:latest
MAINTAINER wangying 

# ENV GOPATH /go

# Install beego & bee
RUN go get github.com/astaxie/beego
RUN go get github.com/beego/bee
RUN go get github.com/go-sql-driver/mysql
