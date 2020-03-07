FROM debian:buster

RUN apt-get update && DEBIAN_FRONTEND=noninteractive \
    apt-get install -y dovecot-imapd dovecot-submissiond dovecot-lmtpd

COPY dovecot/entrypoint.sh /entrypoint.sh

ENTRYPOINT ["/bin/bash", "/entrypoint.sh"]
