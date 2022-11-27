package input

import (
	"fmt"
	"strconv"

	"github.com/Radynsade/typeui/input/rules"
)

// Always adds 'IsFloat' rule on validation so you do not need
// to add it manually. Returns float64 on read.
type Number struct {
	*rules.Set
	BaseInput
}

func NewNumber(
	label string,
	placeholder string,
	required bool,
) *Number {
	rulesSet := rules.NewSet()
	rulesSet.Add(rules.IsFloat())

	return &Number{
		Set: rulesSet,
		BaseInput: BaseInput{
			Label:       label,
			Placeholder: placeholder,
			Required:    required,
		},
	}
}

func (n *Number) SetMax(
	max float64,
	included bool,
) {
	n.Rules = append(n.Rules, rules.MaxNumber(max, included))
}

func (n *Number) SetMin(
	min float64,
	included bool,
) {
	n.Rules = append(n.Rules, rules.MinNumber(min, included))
}

func (n *Number) PrintAndRead() float64 {
	n.PrintLabel()

	var value string

	for {
		n.PrintPlaceholder()

		fmt.Scanln(&value)
		validationMap, allValid := n.Validate(value)

		if allValid {
			number, _ := strconv.ParseFloat(value, 64)

			return number
		}

		if !validationMap["isFloat"] {
			n.PrintError("isFloat")

			continue
		}

		for ruleName, valid := range validationMap {
			if valid || ruleName == "isFloat" {
				continue
			}

			n.PrintError(ruleName)
		}
	}
}
