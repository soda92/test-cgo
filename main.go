package main

/*
#cgo CFLAGS: -g -Wall
#cgo LDFLAGS: -L. -lperson
#include "person.h"
*/
import "C"
import (
	"fmt"
)

type (
	Person C.struct_APerson
)

func (p *Person) Name() string{
	return C.GoString(p.name)
}

func (p *Person) LongName() string{
	return C.GoString(p.long_name)
}

func GetPerson(name string, long_name string) *Person {
	return (*Person)(C.get_person(C.CString(name), C.CString(long_name)))
}

func main() {
	var f *Person = GetPerson("Tim", "Tim Hughes")
	fmt.Printf("Helo Go world: My name is %s, %s.\n",
		C.GoString(f.name), C.GoString(f.long_name))
	fmt.Printf("Helo Go world: My name is %s, %s.\n",
		f.Name(), f.LongName())
}
