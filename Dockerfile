# syntax=docker/dockerfile:experimental
FROM --platform=linux/amd64 golang:1.18 AS build-env
ENV APP_HOME /practice
ENV GOFLAGS -mod=vendor
WORKDIR $APP_HOME
ADD . $APP_HOME
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $APP_HOME/bin/consumer ./cmd/consumer
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $APP_HOME/bin/api ./cmd/api

FROM --platform=linux/amd64 ruby:3.1.2-slim
RUN apt-get update && apt-get install -y build-essential libpq-dev && rm -rf /var/lib/apt/lists/*

ENV BUNDLE_VERSION=2.4.10

WORKDIR /app
COPY db_migrations db_migrations
RUN gem install bundler -v $BUNDLE_VERSION

RUN bundle _${BUNDLE_VERSION}_ install --gemfile db_migrations/Gemfile

COPY run-*.sh ./
COPY --from=build-env /practice/bin/* /app/
CMD ["./run-api.sh"]
