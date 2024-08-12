OPTIONS=-target=wasi -tags=purego

.PHONY: build

build:
	tinygo build $(OPTIONS) -o build/test_app.wasm test_app/main.go
