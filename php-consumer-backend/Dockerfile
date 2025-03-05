FROM php:8.4-fpm-alpine

COPY --from=composer:latest /usr/bin/composer /usr/local/bin/composer

RUN apk add postgresql-dev libffi-dev pact-ffi-lib
RUN docker-php-ext-configure pgsql -with-pgsql=/usr/local/pgsql
RUN docker-php-ext-install opcache pdo pdo_pgsql ffi
RUN docker-php-ext-configure ffi

USER www-data
