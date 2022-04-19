package test_cgo

/*
#cgo CFLAGS: -g -Wall
#cgo LDFLAGS: -L. -lperson
#include "person.h"
*/
import "C"

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
