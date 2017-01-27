FROM ubuntu:trusty
MAINTAINER cdstelly <cdstelly@gmail.com>
RUN apt-get update

RUN apt-get install -y curl exiftool coreutils

ADD bin/rpcserver /
ADD bin/rpcclient /

CMD ["/rpcserver"]
