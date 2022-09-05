FROM golang:1.18.2-buster

ARG TOKEN

RUN apt-get update && apt-get install -y git
ENV GOCACHE=/tmp/build/.cache
ENV GOMODCACHE=/tmp/build/.modcache

COPY . /tmp/tools

RUN cd /tmp \
  && mkdir -p /tmp/build/.cache \
  && mkdir -p /tmp/build/.modcache \
  && cd /tmp/tools \
  && go install -trimpath -tags=tools github.com/golangci/golangci-lint/cmd/golangci-lint \
  && chmod -R 777 /tmp/build/

RUN git config --global url."https://zapci:${TOKEN}@github.com".insteadOf "https://github.com" \
  && go env -w GOPRIVATE=github.com/LN-Zap/btcwallet-strike

WORKDIR /build
