godep:
	go get github.com/tools/godep
	godep restore ./...

setup: godep

test: godep
	godep go test ./...
