package forms

type Field struct {
	Name       string
	Value      string
	Validators []Validator
}

type ValidationResult map[string]bool

func NewField(
	name string,
	validators ...Validator,
) *Field {
	return &Field{
		Name:       name,
		Value:      "",
		Validators: validators,
	}
}

func (f *Field) Validate() ValidationResult {
	result := make(ValidationResult)

	for _, validator := range f.Validators {
		name, valid := validator(f.Value)
		result[name] = valid
	}

	return result
}
