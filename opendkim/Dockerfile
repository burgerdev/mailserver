FROM debian:bookworm

RUN apt-get update && DEBIAN_FRONTEND=noninteractive \
    apt-get install -y opendkim opendkim-tools

RUN mkdir -p /etc/dkimkeys && chmod 700 /etc/dkimkeys

COPY entrypoint.sh /entrypoint.sh
COPY rsyslog.conf /etc/rsyslog.conf

ENTRYPOINT ["/bin/bash", "/entrypoint.sh"]
