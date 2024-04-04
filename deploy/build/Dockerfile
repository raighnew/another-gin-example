ARG REGISTRY=docker.io
FROM ${REGISTRY}/golang:1.18-alpine3.16 AS builder

ARG APP_RELATIVE_PATH

COPY .. /data/app
WORKDIR /data/app

RUN rm -rf /data/app/bin/
RUN go mod tidy && go build -ldflags="-s -w" -o ./bin/server ${APP_RELATIVE_PATH}
RUN mv config /data/app/bin/


FROM ${REGISTRY}/alpine:3.16

RUN apk add tzdata && cp /usr/share/zoneinfo/Asia/Singapore /etc/localtime \
    && echo "Asia/Singapore" > /etc/timezone \
    && apk del tzdata


ARG APP_ENV
ENV APP_ENV=${APP_ENV}

WORKDIR /data/app
COPY --from=builder /data/app/bin /data/app

EXPOSE 8000
ENTRYPOINT [ "./server" ]

#docker build -t  1.1.1.1:5000/demo-api:v1 --build-arg APP_CONF=config/prod.yml --build-arg  APP_RELATIVE_PATH=./cmd/server/...  .
#docker run -it --rm --entrypoint=ash 1.1.1.1:5000/demo-api:v1