package llvm

import (
	"fmt"
	"strings"
)

type Parameter struct {
	Name string
	typ Type
}

func (p Parameter) Type() Type {
	return p.typ
}

func (p Parameter) Identifier() string {
	return "%" + p.Name
}

func NewParameter(name string, typ Type) *Parameter {
	return &Parameter{
		Name: name,
		typ: typ,
	}
}

type Function struct {
	Name string
	ReturnType Type
	Params []*Parameter

	module *Module
	blocks []*block

	ssaCounter int
}

func (m *Module) NewFunction(name string, returnType Type, params ...*Parameter) *Function {
	fn := &Function{
		Name: name,
		ReturnType: returnType,
		Params: params,

		blocks: []*block{},
		module: m,

		ssaCounter: 0,
	}

	m.functions = append(m.functions, fn)

	return fn
}

func (f *Function) nextSSA() string {
	f.ssaCounter++

	return fmt.Sprintf("%%%d", f.ssaCounter)
}

func (f *Function) String() string {
	var o strings.Builder

	// Get parameter operands
	parameterStrings := make([]string, len(f.Params))
	for i, p := range f.Params {
		parameterStrings[i] = operand(p)
	}

	keyword := "declare"
	if len(f.blocks) > 0 {
		keyword = "define"
	}

	// <keyword> <return type> @<function name>(<parameter types>)
	fmt.Fprintf(&o, "%s %s @%s(%s)", keyword, f.ReturnType, mangleName(f.module.Name, f.Name), strings.Join(parameterStrings, ", "))

	// If no blocks (external function), return the string
	if len(f.blocks) == 0 {
		printBlankLine(&o)

		return o.String()
	}

	fmt.Fprintf(&o, " {\n")

	for _, block := range f.blocks {
		fmt.Fprintf(&o, "%s", block)
	}

	fmt.Fprintf(&o, "}\n")

	return o.String()
}
