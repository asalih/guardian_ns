GOFILES = $(shell find . -name '*.go')

default: build

workdir:
	mkdir -p workdir

build: workdir/guardian_ns

build-native: $(GOFILES)
	go build -o workdir/native-guardian_ns .

workdir/guardian_ns: $(GOFILES)
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o workdir/guardian_ns .