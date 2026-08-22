package llvm

type Builder struct {
	context *Context
	insertionPoint *block
}

func NewBuilder(ctx *Context) *Builder {
	return &Builder{
		context: ctx,
		insertionPoint: nil,
	}
}

func (b *Builder) SetInsertionPoint(block *block) {
	b.insertionPoint = block
}

func (b *Builder) UnsetInsertionPoint() {
	b.insertionPoint = nil
}

func (b *Builder) Insert(instruction Instruction) {
	if b.insertionPoint == nil {
		panic("No insertion point set for builder")
	}

	b.insertionPoint.instructions = append(b.insertionPoint.instructions, instruction)
}

func (b *Builder) TerminateBlock() {
	if b.insertionPoint == nil {
		panic("No insertion point set for builder")
	}

	b.insertionPoint.terminate()
}
