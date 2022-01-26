FROM debian:bullseye-20211011
LABEL Vendor="http-replay" \
      maintainer="NineWoranop@users.noreply.github.com"

ADD https://github.com/NineWoranop/http-replay/releases/download/v0.1/http-replay-linux-amd64.tar.gz  /

RUN cd /;tar -xvzf http-replay-linux-amd64.tar.gz;rm -rf http-replay-linux-amd64.tar.gz;

EXPOSE 9191

# start service
ENTRYPOINT ["/http-replay-linux-amd64"]
