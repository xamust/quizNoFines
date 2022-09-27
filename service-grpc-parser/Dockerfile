FROM golang:latest
MAINTAINER Stepan K. <xamust@gmail.com>
WORKDIR /service-grpc-parser/
VOLUME ["/opt/service-grpc-parser"]
COPY . ./
RUN make build
CMD [ "build/service-grpc-parser" ]