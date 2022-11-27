package input

import (
	"fmt"
	"strconv"

	"github.com/Radynsade/typeui/input/rules"
)

// Always adds 'IsFloat' rule on validation so you do not need
// to add it manually. Returns float64 on read.
type Number struct {
	BaseInput
}

func NewNumber(
	label string,
	placeholder string,
	required bool,
) *Number {
	return &Number{
		BaseInput: BaseInput{
			Label:       label,
			Placeholder: placeholder,
			Required:    required,
			Rules:       []rules.Rule{rules.IsFloat()},
			Messages:    make(map[string]string),
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
		validationMap, allValid := rules.Validate(value, n.Rules)

		if allValid {
			number, _ := strconv.ParseFloat(value, 64)

			return number
		}

		if !validationMap["isFloat"] {
			n.printValidationError("isFloat")

			continue
		}

		for ruleName, valid := range validationMap {
			if valid || ruleName == "isFloat" {
				continue
			}

			n.printValidationError(ruleName)
		}
	}
}

func (n *Number) printValidationError(ruleName string) {
	if message, ok := n.Messages[ruleName]; ok {
		println("- " + message)
	} else {
		println("- " + ruleName)
	}
}
