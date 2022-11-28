# Copyright IBM Corp. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0

ARG GO_VER
ARG ALPINE_VER
FROM alpine:${ALPINE_VER} as base
RUN apk add --no-cache tzdata \
        libressl-dev \
        py3-pytest \
        py3-pytest-xdist \
        zip \
        libxslt \
        doxygen \
        graphviz \
        py3-yaml \
        valgrind
# set up nsswitch.conf for Go's "netgo" implementation
# - https://github.com/golang/go/blob/go1.9.1/src/net/conf.go#L194-L275
# - docker run --rm debian:stretch grep '^hosts:' /etc/nsswitch.conf
RUN echo 'hosts: files dns' > /etc/nsswitch.conf

#FROM golang:${GO_VER}-alpine${ALPINE_VER} as golang
#RUN apk add --no-cache \
#	bash \
#	binutils-gold \
#	gcc \
#	git \
#	make \
#	musl-dev
FROM outis291262/golang-pqcrypto:1.0 as golang
ADD . $GOPATH/src/github.com/hyperledger/fabric
WORKDIR $GOPATH/src/github.com/hyperledger/fabric

FROM golang as orderer
ARG GO_TAGS
ENV LD_LIBRARY_PATH /usr/local/lib
RUN make orderer GO_TAGS=${GO_TAGS}

FROM base
ENV FABRIC_CFG_PATH /etc/hyperledger/fabric
VOLUME /etc/hyperledger/fabric
VOLUME /var/hyperledger
ENV LD_LIBRARY_PATH /usr/local/lib
COPY --from=orderer /usr/local/lib /usr/local/lib
COPY --from=orderer /usr/local/go/src/crypto /usr/local/go/src/crypto
COPY --from=orderer /go/src/github.com/hyperledger/fabric/build/bin /usr/local/bin
COPY --from=orderer /go/src/github.com/hyperledger/fabric/sampleconfig/msp ${FABRIC_CFG_PATH}/msp
COPY --from=orderer /go/src/github.com/hyperledger/fabric/sampleconfig/orderer.yaml ${FABRIC_CFG_PATH}
COPY --from=orderer /go/src/github.com/hyperledger/fabric/sampleconfig/configtx.yaml ${FABRIC_CFG_PATH}
EXPOSE 7050
CMD ["orderer"]
