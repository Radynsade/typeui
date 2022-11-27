package input

import (
	"github.com/Radynsade/typeui"
	"github.com/Radynsade/typeui/input/rules"
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
	// Validation rules.
	Rules    []rules.Rule
	Messages map[string]string
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

func (bi *BaseInput) PrintError(ruleName string) {
	if message, ok := bi.Messages[ruleName]; ok {
		println("✗ " + message)
	} else {
		println("✗ " + ruleName)
	}
}

func (bi *BaseInput) AddRule(rule rules.Rule) {
	bi.Rules = append(bi.Rules, rule)
}

func (bi *BaseInput) SetMessage(
	ruleName string,
	message string,
) {
	bi.Messages[ruleName] = message
}
