dep = ${GOPATH}/bin/dep
curDir = $(shell pwd)
vendor = $(curDir)/vendor
target = tradfri

all: $(target)

tradfri: $(dep) $(vendor) cmd cmd/*
	go build -v

$(dep):
	go get -u github.com/golang/dep/cmd/dep

$(vendor):
	dep ensure -v

test: tradfri
	$(target) list

clean:
	rm -rf $(vendor); rm tradfri