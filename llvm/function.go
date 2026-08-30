package llvm

import (
	"fmt"
	"strings"
)

type Parameter struct {
	Name string
	typ  Type
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
		typ:  typ,
	}
}

type Function struct {
	Name       string
	ReturnType Type
	Params     []*Parameter

	module *Module
	blocks []*block

	ssaCounter int
	noMangle   bool
}

type NewFunctionOptions struct {
	NoMangle bool
}

func (m *Module) NewFunction(opts NewFunctionOptions, name string, returnType Type, params ...*Parameter) *Function {
	fn := &Function{
		Name:       name,
		ReturnType: returnType,
		Params:     params,

		blocks: []*block{},
		module: m,

		ssaCounter: 0,
		noMangle:   opts.NoMangle,
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
		str := p.Type().String()
		if len(f.blocks) > 0 {
			str = operand(p)
		}

		parameterStrings[i] = str
	}

	keyword := "declare"
	if len(f.blocks) > 0 {
		keyword = "define"
	}

	name := f.Name
	if !f.noMangle {
		name = mangleName(f.module.Name, f.Name)
	}

	// <keyword> <return type> @<function name>(<parameter types>)
	fmt.Fprintf(&o, "%s %s @%s(%s)", keyword, f.ReturnType, name, strings.Join(parameterStrings, ", "))

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
