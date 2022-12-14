package input

import (
	"fmt"

	"github.com/Radynsade/typeui/input/rules"
)

type Text struct {
	*rules.Set
	BaseInput
}

func NewText(
	label string,
	placeholder string,
	required bool,
) *Text {
	return &Text{
		Set: rules.NewSet(),
		BaseInput: BaseInput{
			Label:       label,
			Placeholder: placeholder,
			Required:    required,
		},
	}
}

func (t *Text) PrintAndRead() string {
	t.PrintLabel()

	var value string

	for {
		t.PrintPlaceholder()

		fmt.Scanln(&value)
		validationMap, allValid := t.Validate(value)

		if allValid {
			return value
		}

		for ruleName, valid := range validationMap {
			if valid {
				continue
			}

			t.PrintError(ruleName)
		}
	}
}
