FROM centos:latest
RUN yum -y update && yum -y clean all
RUN yum -y install git epel-release
RUN yum -y install --enablerepo=epel rrdtool-devel rrdtool golang

ENV GOPATH /go
ADD . /go/src/app
WORKDIR /go/src/app
RUN go get -u github.com/ziutek/rrd
RUN go build

CMD ["./app"]
