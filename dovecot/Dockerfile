FROM debian:buster

RUN apt-get update && DEBIAN_FRONTEND=noninteractive \
    apt-get install -y dovecot-imapd dovecot-submissiond dovecot-lmtpd

RUN useradd -M -U -s /usr/sbin/nologin -d /vmail -u 5000 vmail

COPY dovecot/entrypoint.sh /entrypoint.sh

ENTRYPOINT ["/bin/bash", "/entrypoint.sh"]
