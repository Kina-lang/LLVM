package llvm

import "martinpetr.dev/kina/llvm/llvmtarget"

type Context struct {
	target llvmtarget.Target
}

func NewContext(triple string) (*Context, error) {
	target, err := llvmtarget.NewTarget(llvmtarget.Triple(triple))
	if err != nil {
		return nil, err
	}

	return &Context{
		target: *target,
	}, nil
}
