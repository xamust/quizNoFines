FROM golang:latest
MAINTAINER Stepan K. <xamust@gmail.com>
WORKDIR /service-grpc-gw/
VOLUME ["/opt/service-grpc-gw"]
COPY . ./
RUN make build
CMD [ "build/service-grpc-gw" ]