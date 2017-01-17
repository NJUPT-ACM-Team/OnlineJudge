export GOPATH=$(shell pwd)

default: all

deps:
	go get -d -v OnlineJudge/... Daemon/... VJudger/... WebBackend/...

WebBackend: deps
	go install WebBackend

Daemon: deps
	go install Daemon

VJudger: deps
	go install VJudger

fmt:
	go fmt OnlineJudge/... Daemon/... VJudger/... WebBackend/...

all: fmt WebBackend Daemon VJudger
