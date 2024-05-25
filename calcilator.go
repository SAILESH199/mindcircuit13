FROM amazonlinux
RUN yum update && yum install -y go
ENV GO111MODULE=off
WORKDIR /opt
COPY . .
RUN go build calculator.go
ENTRYPOINT ["/opt/calculator"]
[root@docker calculatorapp]#
