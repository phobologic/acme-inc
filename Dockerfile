FROM golang:1.7.1-alpine
MAINTAINER Michael Barrett <loki77@gmail.com>

# Install ssm-env cleanly
RUN apk update && \
    apk add curl && \
    curl -L https://github.com/remind101/ssm-env/releases/download/v0.0.4/ssm-env > /usr/local/bin/ssm-env && \
    cd /usr/local/bin && \
    echo "4a5140b04f8b3f84d16a93540daa7bbd  ssm-env" | md5sum -c && \
    chmod +x ssm-env && \
    apk del curl

COPY ./ /go/src/github.com/phobologic/acme-inc
RUN go install github.com/phobologic/acme-inc
WORKDIR /go/src/github.com/phobologic/acme-inc

ENTRYPOINT ["/usr/local/bin/ssm-env", "-with-decryption"]
CMD ["acme-inc", "server"]
