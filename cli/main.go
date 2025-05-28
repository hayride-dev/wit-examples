package main

import (
	"github.com/hayride-dev/wit-examples/internal/imports/hayride/wit-examples/bar"
	"github.com/hayride-dev/wit-examples/internal/imports/hayride/wit-examples/foobar"
)

// CLI
func init() {
}

func main() {
	foo := foobar.NewFoo()
	bar := bar.NewBar()

	value := foobar.Foobar(bar, foo)
	println("Value from foobar:", value)
}
