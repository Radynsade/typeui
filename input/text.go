package input

import (
	"fmt"

	"github.com/Radynsade/typeui/input/rules"
)

type Text struct {
	BaseInput
}

func NewText(
	label string,
	placeholder string,
	required bool,
	rules ...rules.Rule,
) *Text {
	return &Text{
		BaseInput: BaseInput{
			Label:       label,
			Placeholder: placeholder,
			Required:    required,
			Rules:       rules,
			Messages:    make(map[string]string),
		},
	}
}

func (t *Text) PrintAndRead() string {
	t.PrintLabel()

	var value string

	for {
		t.PrintPlaceholder()

		fmt.Scanln(&value)
		validationMap, allValid := rules.Validate(value, t.Rules)

		if allValid {
			return value
		}

		for ruleName, valid := range validationMap {
			if valid {
				continue
			}

			t.printValidationError(ruleName)
		}
	}
}

func (t *Text) printValidationError(ruleName string) {
	if message, ok := t.Messages[ruleName]; ok {
		println("- " + message)
	} else {
		println("- " + ruleName)
	}
}
