package llvm

import (
	"strconv"
)

type Value interface {
	Type() Type
	Identifier() string
}

type Const struct {
	typ Type
	value int64
}

func (c Const) Type() Type {
	return c.typ
}

func (c Const) Identifier() string {
	return strconv.FormatInt(c.value, 10)
}

func (t IntType) Const(n int64) Const {
	return Const{t, n}
}

type VoidConst struct{}
func (VoidConst) Type() Type {
	return Void
}

func (VoidConst) Identifier() string {
	return ""
}

func (voidType) Const() VoidConst {
	return VoidConst{}
}

// <type> <ident>
func operand(v Value) string {
	return v.Type().String() + " " + v.Identifier()
}
