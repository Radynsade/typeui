package input

import (
	"fmt"

	"github.com/Radynsade/typeui/input/rules"
)

type YesNo struct {
	*rules.Set
	BaseInput
}

func NewYesNo(
	label string,
	placeholder string,
	required bool,
) *YesNo {
	rulesSet := rules.NewSet()
	rulesSet.Add(rules.OneOf("y", "n"))

	return &YesNo{
		Set: rulesSet,
		BaseInput: BaseInput{
			Label:       label,
			Placeholder: placeholder,
			Required:    required,
		},
	}
}

func (yn *YesNo) PrintAndRead() bool {
	yn.PrintLabel()

	var value string

	for {
		yn.PrintPlaceholder()

		fmt.Scanln(&value)
		validationMap, allValid := yn.Validate(value)

		if allValid {
			if value == "y" {
				return true
			}

			return false
		}

		for ruleName, valid := range validationMap {
			if valid {
				continue
			}

			yn.PrintError(ruleName)
		}
	}
}
