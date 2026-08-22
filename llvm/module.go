package llvm

import (
	"fmt"
	"strings"

	"martinpetr.dev/kina/llvm/llvmtarget"
)

type Module struct {
	Name    string
	Context *Context
	functions []*Function
	aliases  []*functionAlias
}

func NewModule(name string, context *Context) *Module {
	return &Module{
		Name:    name,
		Context: context,
	}
}

func (m *Module) String() string {
	var o strings.Builder

	printMeta(&o, m)
	printBlankLine(&o)

	for _, fn := range m.functions {
		printComment(&o, fmt.Sprintf("Function %s", fn.Name))
		fmt.Fprintf(&o, "%s\n", fn.String())
	}

	for _, alias := range m.aliases {
		printComment(&o, fmt.Sprintf("Function alias %s (target: %s)", alias.Name, alias.Function.Name))
		fmt.Fprintf(&o, "%s\n", alias.String())
	}

	return o.String()
}

func printTargetMeta(o *strings.Builder, target *llvmtarget.Target) {
	fmt.Fprintf(o, "target datalayout = %q\n", target.Layout)
	fmt.Fprintf(o, "target triple = %q\n", target.Triple)
}

func printMeta(o *strings.Builder, m *Module) {
	printComment(o, fmt.Sprintf("ModuleID = '%s'", m.Name))
	fmt.Fprintf(o, "source_filename = %q\n", m.Name)

	printTargetMeta(o, &m.Context.target)
}

type functionAlias struct {
	Name string
	Function *Function
}

func (a *functionAlias) String() string {
	var params []string
	for _, p := range a.Function.Params {
		params = append(params, p.Type().String())
	}

	return fmt.Sprintf("@%s = alias %s, ptr @%s", a.Name, fmt.Sprintf("%s (%s)", a.Function.ReturnType, strings.Join(params, " ")), mangleName(a.Function.module.Name, a.Function.Name))
}

func (m *Module) NewFunctionAlias(name string, function *Function) *functionAlias {
	a := &functionAlias{
		Name: name,
		Function: function,
	}

	m.aliases = append(m.aliases, a)

	return a
}
