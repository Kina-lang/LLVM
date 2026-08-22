package llvm

type Instruction interface {
	String() string
}

// Return instruction (`ret <type> <value>` or `ret void`)
type ret struct{
	value Value
}

func (r ret) String() string {
	if r.value == nil {
		return "ret void"
	}

	return "ret " + operand(r.value)
}

// Return instruction
// Terminates the current block
func (b *Builder) CreateRet(val Value) {
	i := &ret{val}
	b.Insert(i)
	b.TerminateBlock()
}
