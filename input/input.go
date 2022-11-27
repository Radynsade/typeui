package input

import (
	"github.com/Radynsade/typeui"
)

type Input[ValueType comparable] interface {
	typeui.Node
	PrintLabel()
	PrintPlaceholder()
	PrintAndRead() ValueType
}

// DOES NOT IMPLEMENT 'Input' INTERFACE COMPLETELY!
type BaseInput struct {
	// Prints on separate line before placeholder. Usually means a name of the input.
	// If empty then 'PrintLabel' will do nothing.
	Label string
	// Prints before cursor on the same line. Cannot be empty.
	Placeholder string
	// Adds '*' to the end of the placeholder if true.
	Required bool
}

func (bi *BaseInput) PrintLabel() {
	if len(bi.Label) <= 0 {
		return
	}

	println(bi.Label)
}

func (bi *BaseInput) PrintPlaceholder() {
	if bi.Required {
		print(bi.Placeholder + "*: ")
	} else {
		print(bi.Placeholder + ": ")
	}
}

func (bi *BaseInput) Print() {
	bi.PrintLabel()
	bi.PrintPlaceholder()
}
