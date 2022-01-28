FROM golang:1.17.5-alpine3.15 AS builder
LABEL version="1.17.5"
RUN apk --update add git git-lfs less openssh && \
    git lfs install && \
    rm -rf /var/lib/apt/lists/* && \
    rm /var/cache/apk/*
RUN mkdir /builder
WORKDIR /builder
COPY . /builder
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -a -o app .

FROM scratch
MAINTAINER fmessina
LABEL version="0.0.1"
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /builder/app .
USER 1001
ENTRYPOINT ["./app"]
EXPOSE 8080
