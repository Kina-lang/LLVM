package llvmtarget

import "fmt"

type Target struct {
	Triple Triple
	Layout string
}

func NewTarget(triple Triple) (*Target, error) {
	layout, ok := Layouts[triple]
	if !ok {
		return nil, fmt.Errorf("llvm: unsupported target triple %q", triple)
	}

	return &Target{
		Triple: triple,
		Layout: layout,
	}, nil
}

type Triple string

const (
	X8664_Unknown_Linux_Gnu Triple = "x86_64-unknown-linux-gnu"
)

// LLVM datalayout map for each triple
// This can be obtained by running:
// clang -target <TARGET> -S -emit-llvm -x c /dev/null -o - | grep '^target datalayout'
var Layouts = map[Triple]string{
	X8664_Unknown_Linux_Gnu: "e-m:e-p270:32:32-p271:32:32-p272:64:64-i64:64-i128:128-f80:128-n8:16:32:64-S128",
}
