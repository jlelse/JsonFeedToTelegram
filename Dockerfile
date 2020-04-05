FROM golang:1.14-alpine as build
ADD . /app
WORKDIR /app
RUN go build

FROM alpine:3.11
RUN apk add --no-cache tzdata ca-certificates
COPY --from=build /app/JsonFeedToTelegram /bin/
WORKDIR /app
VOLUME /app/storage
EXPOSE 8080
ENV LAST_ARTICLE_FILE /app/storage/lastarticle
CMD ["JsonFeedToTelegram"]