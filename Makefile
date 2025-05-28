.PHONY: gen-ml gen-http gen-morph gen-silo gen-core gen

default: gen

gen-exports:
	wit-bindgen-go generate -v -vv --world hayride:wit-examples/exports --out ./internal ./wit

gen-imports:
	wit-bindgen-go generate -v -vv --world hayride:wit-examples/imports --out ./internal/imports ./wit

gen: gen-exports gen-imports
