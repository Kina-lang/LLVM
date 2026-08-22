package llvm

import (
	"fmt"
	"strings"
)

type block struct {
	label string
	instructions []Instruction
	terminated bool
}

func (f *function) NewBlock(label string) *block {
	block := &block{
		label: label,
		instructions: []Instruction{},
		terminated: false,
	}

	f.blocks = append(f.blocks, block)

	return block
}

func (b *block) String() string {
	var o strings.Builder

	fmt.Fprintf(&o, "%s:\n", b.label)

	for _, instruction := range b.instructions {
		fmt.Fprintf(&o, "%s\n", indentLine(instruction.String(), 1))
	}

	return o.String()
}

func (b *block) terminate() {
	if b.terminated {
		panic("Block already terminated")
	}

	b.terminated = true
}
