package forms

type Field struct {
	Name  string
	Value string
	Rules []Rule
}

type ValidationResult map[string]bool

func NewField(
	name string,
	rules ...Rule,
) *Field {
	return &Field{
		Name:  name,
		Value: "",
		Rules: rules,
	}
}

func (f *Field) SetValue(value string) {
	f.Value = value
}

func (f *Field) Validate() ValidationResult {
	result := make(ValidationResult)

	for _, validator := range f.Rules {
		name, valid := validator(f.Value)
		result[name] = valid
	}

	return result
}
