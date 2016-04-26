GOC=go build
GOFLAGS=-a -ldflags '-s'
CGOR=CGO_ENABLED=0

all: client

build:
	$(GOC) client.go

run:
	go run client.go

stat:
	$(CGOR) $(GOC) $(GOFLAGS) client.go

install: stat
	cp client /usr/bin

fmt:
	gofmt -w .

clean:
	rm client

#CGO_ENABLED=0 go build -a -ldflags '-s' client.go
