FROM alpine:latest AS builder

WORKDIR /var/www/html

RUN wget -O rainloop.zip https://github.com/RainLoop/rainloop-webmail/releases/download/v1.17.0/rainloop-legacy-1.17.0.zip && \
    echo '782dcabacadab5d7176f7701dd23319a040b2cfbf974fac6df068600cf69c50a rainloop.zip' | sha256sum -c && \
    unzip rainloop.zip && \
    rm -f rainloop.zip && \
    find . -type d -exec chmod 755 {} \; && \
    find . -type f -exec chmod 644 {} \; && \
    chown -R 33:33 . && \
    printf "order allow,deny\ndeny from all\n" >/var/www/html/data/.htaccess

FROM php:apache-bookworm

RUN mv /usr/local/etc/php/php.ini-production /usr/local/etc/php/php.ini

COPY --from=builder /var/www/html /var/www/html
