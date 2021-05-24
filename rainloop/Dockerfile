FROM php:7.4-apache

RUN apt-get update && \
    DEBIAN_FRONTEND=noninteractive apt-get install -y unzip && \
    mv /usr/local/etc/php/php.ini-production /usr/local/etc/php/php.ini && \
    cd /var/www/html && \
    curl -L https://github.com/RainLoop/rainloop-webmail/releases/download/v1.16.0/rainloop-1.16.0.zip >rainloop.zip && \
    echo '6980b0a85aa124479a927e11d0c50319c4a9f8ce3b25b6c2adb61a30bf1a792b  rainloop.zip' | sha256sum -c && \
    unzip rainloop.zip && \
    rm -f rainloop.zip && \
    apt-get remove --purge -y unzip && \
    rm -rf /var/lib/apt/lists/* && \
    find . -type d -exec chmod 755 {} \; && \
    find . -type f -exec chmod 644 {} \; && \
    chown -R www-data:www-data . && \
    printf "order allow,deny\ndeny from all\n" >/var/www/html/data/.htaccess
