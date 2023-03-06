FROM php:7.4-apache

RUN apt-get update && \
    DEBIAN_FRONTEND=noninteractive apt-get install -y unzip && \
    mv /usr/local/etc/php/php.ini-production /usr/local/etc/php/php.ini && \
    cd /var/www/html && \
    curl -L https://github.com/RainLoop/rainloop-webmail/releases/download/v1.17.0/rainloop-legacy-1.17.0.zip>rainloop.zip && \
    echo '782dcabacadab5d7176f7701dd23319a040b2cfbf974fac6df068600cf69c50a rainloop.zip' | sha256sum -c && \
    unzip rainloop.zip && \
    rm -f rainloop.zip && \
    apt-get remove --purge -y unzip && \
    rm -rf /var/lib/apt/lists/* && \
    find . -type d -exec chmod 755 {} \; && \
    find . -type f -exec chmod 644 {} \; && \
    chown -R www-data:www-data . && \
    printf "order allow,deny\ndeny from all\n" >/var/www/html/data/.htaccess
