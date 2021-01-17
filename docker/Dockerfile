FROM debian:stable-slim

ADD tmp/simple-server.linux /simple-server

USER 1000

RUN env
ENTRYPOINT ["/simple-server"]
