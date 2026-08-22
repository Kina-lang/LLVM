package llvm

import (
	"fmt"
	"strings"

	"martinpetr.dev/kina/llvm/llvmtarget"
)

type Module struct {
	Name    string
	Context *Context
	functions []*function
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
