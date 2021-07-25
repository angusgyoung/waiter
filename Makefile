BINARY = waiter
VERSION := $(shell git describe --always --long --dirty)
BUILD := $(shell date +%FT%T%z)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/)
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)

LDFLAGS=-ldflags "-w -s -X main.Version=${VERSION} -X main.Build=${BUILD}"

build:
	go build ${LDFLAGS} -o ${BINARY}

install:
	go install ${LDFLAGS}

test:
	go test -short ${PKG_LIST}

vet:
	go vet ${PKG_LIST}

lint:
	for file in ${GO_FILES} ;  do \
		golint $$file ; \
	done

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

.PHONY: clean install


