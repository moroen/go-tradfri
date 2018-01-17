dep = ${GOPATH}/bin/dep
curDir = $(shell pwd)
vendor = $(curDir)/vendor

all: tradfri

tradfri: $(dep) $(vendor)
	go build -v

$(dep):
	go get -u github.com/golang/dep/cmd/dep

$(vendor):
	dep ensure -v

clean:
	rm -rf $(vendor); rm tradfri