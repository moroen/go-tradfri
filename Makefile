dep = ${GOPATH}/bin/dep
curDir = $(shell pwd)
# vendor = $(curDir)/vendor
go-coap-lib = ${GOPATH}/src/github.com/moroen/go-tradfricoap/
go-canopus = ${GOPATH}/src/github.com/moroen/canopus/

target = tradfri

all: $(target)

tradfri: $(dep) $(vendor) $(go-coap-lib)/*.go $(go-canopus)/*.go cmd cmd/* *.go
	go build -v

$(dep):
	go get -u github.com/golang/dep/cmd/dep

$(vendor):
	dep ensure -v

test: tradfri
	./$(target) observe 65545

install: $(target)
	go install
	
clean:
	rm -rf $(vendor); rm -rf $(target)