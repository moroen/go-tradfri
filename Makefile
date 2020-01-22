ifndef GOPATH
$(error GOPATH is not set)
endif

dep = ${GOPATH}/bin/dep
curDir = $(shell pwd)
dependencies = $(GOPATH)/src/github.com/spf13/cobra
files = *.go cmd/*.go

#go-coap-lib = ${GOPATH}/src/github.com/moroen/go-tradfricoap/
#go-canopus = ${GOPATH}/src/github.com/moroen/canopus/

target = tradfri

all: $(target)

tradfri: $(dependencies) $(files)
	go build -v -o tradfri main.go

$(dependencies):
	go get -v

test: tradfri
	./$(target) observe 65545

install: $(target)
	go install
	
clean:
	rm -rf $(vendor); rm -rf $(target)