package typeui

import (
	"bufio"
	"os"
	"strings"

	"github.com/Radynsade/typeui/forms"
)

type Input interface {
	Node
	PrintLabel()
	PrintPlaceholder()
	PrintAndRead()
}

type BaseInput struct {
	// Prints '*' after label. Does not allow empty input.
	Required bool
	// Accept multiple values.
	Multiple bool
	// Delimiter of multiple values. Has no impact if 'Multiple' is false.
	Delimiter string
	// Prints on separate line before placeholder. Usually means a name of the input.
	Label string
	// Prints before cursor on the same line.
	Placeholder string
	// Field to store value in.
	Field *forms.Field
}

func NewBaseInput(
	label string,
	placeholder string,
	field *forms.Field,
	required bool,
) *BaseInput {
	return &BaseInput{
		Label:       label,
		Placeholder: placeholder,
		Field:       field,
		Required:    required,
	}
}

func (i *BaseInput) PrintLabel() {
	if len(i.Label) <= 0 {
		return
	}

	if i.Required {
		println(i.Label + "*")
	} else {
		println(i.Label)
	}
}

func (i *BaseInput) PrintPlaceholder() {
	print(i.Placeholder + ": ")
}

func (i *BaseInput) Print() {
	i.PrintLabel()
	i.PrintPlaceholder()
}

func (i *BaseInput) PrintAndRead() {
	i.PrintLabel()

	for {
		i.PrintPlaceholder()

		value := readLine()

		if i.Field == nil {
			break
		}

		i.Field.SetValue(value)

		result := i.Field.Validate()
		allValid := true

		for rule, valid := range result {
			if !valid {
				allValid = false
				println(rule)
			}
		}

		if allValid {
			break
		}
	}
}

func readLine() string {
	reader := bufio.NewReader(os.Stdin)
	value, _ := reader.ReadString('\n')
	// convert CRLF to LF
	value = strings.Replace(value, "\n", "", -1)

	return value
}

// func readLines() []string {
// 	reader := bufio.NewReader(os.Stdin)
// 	values := make([]string, 0)

// 	for {
// 		value, _ := reader.ReadString('\n')
// 		// convert CRLF to LF
// 		value = strings.Replace(value, "\n", "", -1)
// 		values = append(values, value)
// 	}

// 	return values
// }
