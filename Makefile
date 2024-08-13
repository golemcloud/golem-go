OPTIONS=-target=wasi -tags=purego

.PHONY: default update-wit wit-bindgen build update-and-build

default: all

wit-deps-update:
	wit-deps update

wit-bindgen:
	rm -rf binding
	wit-bindgen tiny-go --out-dir binding --rename-package binding --gofmt ./wit

build:
	tinygo build $(OPTIONS) -o build/test_app.wasm test_app/main.go

all: wit-deps-update wit-bindgen build
