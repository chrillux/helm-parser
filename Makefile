.PHONY: build buildfiles clean prodbuild runls

clean:
	rm -rf build

build:
	GOARCH=wasm GOOS=js go build -o lib.wasm wasm/wasm.go

buildfiles:
	mkdir -p build && cp -r *.wasm css images js index.html build/

prodbuild: clean build buildfiles

run:
	go run main.go
