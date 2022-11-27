package input

type Select[ValueType any] struct {
	BaseInput
	Multiple bool
	MaxItems int
	MinItems int
	Items    []*SelectItem[ValueType]
}

func NewSelect[ValueType any](
	label string,
	placeholder string,
	required bool,
	multiple bool,
) *Select[ValueType] {
	return &Select[ValueType]{
		BaseInput: BaseInput{
			Label:       label,
			Placeholder: placeholder,
			Required:    required,
		},
		Multiple: multiple,
		MaxItems: 0,
		MinItems: 0,
		Items:    make([]*SelectItem[ValueType], 0),
	}
}

type SelectItem[ValueType any] struct {
	Value ValueType
	Label string
}

func NewSelectItem[ValueType any](
	value ValueType,
	label string,
) *SelectItem[ValueType] {
	return &SelectItem[ValueType]{
		Value: value,
		Label: label,
	}
}

func (s *Select[ValueType]) SetMax(max int) {
	if s.Multiple && s.MinItems < max {
		s.MaxItems = max
	}
}

func (s *Select[ValueType]) SetMin(min int) {
	if s.Multiple && s.MaxItems > min {
		s.MinItems = min
	}
}

func (s *Select[ValueType]) AddItem(
	value ValueType,
	label string,
) {
	s.Items = append(s.Items, NewSelectItem(value, label))
}
