OPTIONS=-target=wasi -tags=purego

.PHONY: test

test:
	tinygo test $(OPTIONS) ./...
