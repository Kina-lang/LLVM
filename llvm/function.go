package llvm

import (
	"fmt"
	"strings"
)

type parameter struct {
	Name string
	typ Type
}

func (p parameter) Type() Type {
	return p.typ
}

func (p parameter) Identifier() string {
	return "%" + p.Name
}

type function struct {
	Name string
	ReturnType Type
	Params []parameter

	module *Module
	blocks []*block

	ssaCounter int
}

func (m *Module) NewFunction(name string, returnType Type, params ...parameter) *function {
	fn := &function{
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

func (f *function) nextSSA() string {
	f.ssaCounter++

	return fmt.Sprintf("%%%d", f.ssaCounter)
}

func (f *function) String() string {
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
