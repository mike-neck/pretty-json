.PHONY: build
build:
	go build -o build/pretty-json ./...

.PHONY: clean
clean:
	rm -rf build/
