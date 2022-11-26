package validation

type Field struct {
	Name  string
	Rules []Rule
}

func NewField(
	name string,
	rules ...Rule,
) *Field {
	return &Field{
		Name:  name,
		Rules: rules,
	}
}

func (f *Field) Validate() {
	for _, rule := range f.Rules {

	}
}
