.PHONY: gen-ml gen-http gen-morph gen-silo gen-core gen

default: gen build wac run

gen-exports:
	wit-bindgen-go generate -v -vv --world hayride:wit-examples/exports --out ./internal ./wit

gen-imports:
	wit-bindgen-go generate -v -vv --world hayride:wit-examples/imports --out ./internal/imports ./wit

gen: gen-exports gen-imports

build-cli:
	tinygo build -o cli.wasm -target=wasip2 -wit-package ./wit/ --wit-world imports ./cli

build-foobar:
	tinygo build -o foobar.wasm -target=wasip2 -wit-package ./wit/ --wit-world exports ./foobar

build: build-cli build-foobar

wac: 
	wac plug cli.wasm --plug foobar.wasm -o plugged.wasm

run:
	wasmtime plugged.wasm
