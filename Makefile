.PHONY: build/% run/% compile/%

build/%:
	go build -o bin/$@ $@.go

run/%:
	go run $@.go

compile/%:
	GOOS=darwin GOARCH=amd64 go build -o bin/$@ $@n.go
