FROM debian:stable-slim

RUN apt-get update && apt-get install -y curl jq lsof procps sudo tcpdump wget zip

ADD tmp/simple-server.linux /simple-server

USER 1000

RUN env
ENTRYPOINT ["/simple-server"]
