OPTIONS=-target=wasi -tags=purego

.PHONY: default update-wit wit-bindgen build update-and-build

default: all

wit-deps-update:
	rm -rf wit/deps
	wit-deps update

wit-bindgen:
	rm -rf binding
	wit-bindgen-go -v generate --world golem-go-bindings --out binding ./wit

build:
	tinygo build $(OPTIONS) -o build/test_app.wasm test_app/main.go

all: wit-deps-update wit-bindgen build

# For testing components using the library locally with module replace.
# In the published version binding/binding.c is git ignored, but locally we have to delete it.
prepare-for-local-replace:
	rm binding/binding.c
