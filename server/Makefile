GOC=go build
GOFLAGS=-a -ldflags '-s'
CGOR=CGO_ENABLED=0

all: server

build:
	$(GOC) server.go

run:
	go run server.go

stat:
	$(CGOR) $(GOC) $(GOFLAGS) server.go

install: stat
	cp server /usr/bin

fmt:
	gofmt -w .

clean:
	rm server

#CGO_ENABLED=0 go build -a -ldflags '-s' server.go
