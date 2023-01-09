# Copyright IBM Corp. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0

ARG GO_VER
ARG ALPINE_VER
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

FROM golang as tools
ARG GO_TAGS
RUN make tools GO_TAGS=${GO_TAGS}

FROM outis291262/golang-pqcrypto:1.0
# git is required to support `go list -m`
RUN apk add --no-cache \
	bash \
	git \
	jq \
	tzdata;
ENV FABRIC_CFG_PATH /etc/hyperledger/fabric
ENV LD_LIBRARY_PATH /usr/local/lib
VOLUME /etc/hyperledger/fabric
COPY --from=tools /usr/local/go/src/crypto /usr/local/go/src/crypto
COPY --from=tools /go/src/github.com/hyperledger/fabric/build/bin /usr/local/bin
#COPY --from=tools /go/src/github.com/hyperledger/fabric/sampleconfig ${FABRIC_CFG_PATH}
