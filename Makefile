OPTIONS=-target=wasi -tags=purego

.PHONY: default update-wit wit-bindgen build update-and-build

default: all

wit-deps-update:
	rm -rf wit/deps
	wit-deps update

wit-bindgen:
	rm -rf binding
	go tool wit-bindgen-go -v generate --world golem-go-bindings --out binding ./wit

build:
	tinygo build $(OPTIONS) -o build/test_app.wasm test_app/main.go

all: wit-deps-update wit-bindgen build

