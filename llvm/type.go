package llvm

import (
	"fmt"
	"strconv"
)

type Type interface {
	String() string
}

type IntType struct {
	Bits int
}
func (t IntType) String() string {
	return "i" + strconv.Itoa(t.Bits) // i<bits>
}

type voidType struct {}
func (t voidType) String() string {
	return "void"
}

type ptrType struct{}
func (ptrType) String() string {
	return "ptr" // LLVM opaque pointer type
}

type arrayType struct {
	Count int
	ElemType Type
}
func (t arrayType) String() string {
	return fmt.Sprintf("[%dx%s]", t.Count, t.ElemType)
}

var (
	Int1 = IntType{1}
	Int8 = IntType{8}
	Int16 = IntType{16}
	Int32 = IntType{32}
	Int64 = IntType{64}

	Void = voidType{}

	Ptr = ptrType{}
)
