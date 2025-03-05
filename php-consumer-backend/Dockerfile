FROM php:8.4-fpm-alpine

COPY --from=composer:latest /usr/bin/composer /usr/local/bin/composer

RUN apk add libffi-dev
RUN docker-php-ext-install opcache ffi
RUN docker-php-ext-configure ffi

USER www-data
