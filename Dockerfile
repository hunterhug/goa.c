FROM hunterhug/gitbook:latest AS gitbookk
WORKDIR /srv/gitbook
COPY . /srv/gitbook/
RUN gitbook build .  --log=debug --debug

## main build
FROM nginx:1.15 AS prod
WORKDIR /usr/share/nginx/html
COPY --from=gitbookk  /srv/gitbook/_book /usr/share/nginx/html