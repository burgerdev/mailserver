FROM debian:buster

RUN apt-get update && DEBIAN_FRONTEND=noninteractive \
    apt-get install -y opendkim opendkim-tools

COPY opendkim/entrypoint.sh /entrypoint.sh

ENTRYPOINT ["/bin/bash", "/entrypoint.sh"]
