FROM golang:1.20-alpine

ENV GOAPP sean-home
ENV GOPATH /go
ENV USER sean

RUN apk update && \
    apk add --no-cache \
        git \
        gcc \
        libc-dev \
        dumb-init && \
    adduser -D ${USER} && \
    # mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2 && \
    mkdir -p ${GOPATH}/src ${GOPATH}/bin

COPY . ${GOPATH}/src/${GOAPP}/
RUN cd ${GOPATH}/src/${GOAPP}/ && go mod vendor
RUN cd ${GOPATH}/src/${GOAPP} && go build -o $GOAPP
# RUN CGO_ENABLED=1 CC=aarch64-linux-gnu-gcc GOOS=linux GOARCH=arm64 go build -o goSite .
USER ${USER}
WORKDIR ${GOPATH}/src/${GOAPP}
# ENTRYPOINT ["/usr/bin/dumb-init", "--", "/go/src/hockeyTrainer/hockeyTrainer"]
CMD ["/bin/sh"]
