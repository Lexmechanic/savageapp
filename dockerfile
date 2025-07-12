FROM node:current-alpine3.22 AS build_frontend
COPY frontend ./frontend
WORKDIR /frontend
RUN npm install && npm run build

FROM golang:1.24.2-alpine AS build_backend
COPY . /go/src/app
WORKDIR /go/src/app
RUN go build -o savageapp

FROM node:current-alpine3.22
COPY --from=build_frontend frontend/build /app
WORKDIR /app
COPY --from=build_backend /go/src/app/savageapp /app/savageapp
COPY --from=build_backend /go/src/app/db/migrations /app/db/migrations
RUN mkdir -p /app/sqlite

RUN apk add --no-cache nginx
COPY nginx_reverse_proxy.conf /etc/nginx/nginx.conf

RUN apk add --no-cache supervisor
COPY supervisord.conf /etc/supervisord.conf

EXPOSE 80
CMD ["/usr/bin/supervisord", "-c", "/etc/supervisord.conf"]
