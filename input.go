package typeui

type Input[ValueType any] struct {
	Key         string
	Label       string
	Placeholder string
}

func (i *Input[ValueType]) AddRules() {

}

func (i *Input[ValueType]) Read() (ValueType, error) {
	var result ValueType

	println(i.Label)

	return result, nil
}
