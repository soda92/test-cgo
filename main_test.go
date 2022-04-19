package test_cgo

import (
	"testing"
	"fmt"
)

func TestCgo(t *testing.T) {
    var f *Person = GetPerson("Tim", "Tim Hughes")
	fmt.Printf("Helo Go world: My name is %s, %s.\n",
		f.Name(), f.LongName())
}