FROM debian:bookworm

RUN apt-get update && DEBIAN_FRONTEND=noninteractive \
    apt-get install -y postfix

COPY postfix/entrypoint.sh /entrypoint.sh

ENTRYPOINT ["/bin/bash", "/entrypoint.sh"]
