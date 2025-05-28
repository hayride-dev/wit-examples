.PHONY: gen-ml gen-http gen-morph gen-silo gen-core gen

default: gen

gen-exports:
	wit-bindgen-go generate --world hayride:wit-examples/exports --out ./internal ./wit

gen: gen-exports