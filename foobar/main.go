package main

import (
	"fmt"

	"github.com/hayride-dev/wit-examples/internal/hayride/wit-examples/bar"
	"github.com/hayride-dev/wit-examples/internal/hayride/wit-examples/foobar"
	"go.bytecodealliance.org/cm"
)

// EXPORTED FOOBAR
func init() {
	foobar.Exports.Foo.Constructor = fooConstructor
	foobar.Exports.Foo.Fun = fooFunc

	bar.Exports.Bar.Constructor = barConstructor

	foobar.Exports.Foobar = FoobarFunction
}

type fooStruct struct {
	name string
}

func fooFunc(self cm.Rep) (result string) {
	foo := fooTable[uint32(self)]

	fmt.Println("Got foo:", foo.name)

	fmt.Println("Foo method called, self:", self)
	return "Hello from Foo"
}

func fooConstructor() foobar.Foo {
	value := foobar.FooResourceNew(23)

	fooTable[23] = fooStruct{}

	return value
}

func barConstructor() bar.Bar {
	value := bar.BarResourceNew(42)
	return value
}

var fooTable = make(map[uint32]fooStruct)

func FoobarFunction(bar foobar.Bar, f foobar.Foo) string {

	barRep := bar.ResourceRep()
	fooRep := f.ResourceRep()

	fmt.Println("Foobar function called with bar:", barRep, "and foo:", fooRep)

	value := fooFunc(cm.Rep(f))

	return value
}

func main() {

}
