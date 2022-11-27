package rules

type Set struct {
	Rules    []Rule
	Messages map[string]string
}

func NewSet() *Set {
	return &Set{
		Rules:    make([]Rule, 0),
		Messages: make(map[string]string),
	}
}

func (s *Set) Add(rule Rule) {
	s.Rules = append(s.Rules, rule)
}

func (s *Set) SetMessage(
	ruleName string,
	message string,
) {
	s.Messages[ruleName] = message
}

func (s *Set) Validate(value string) (
	results map[string]bool,
	allValid bool,
) {
	results = make(map[string]bool)
	allValid = true

	for _, rule := range s.Rules {
		ruleName, valid := rule(value)
		results[ruleName] = valid

		if !valid {
			allValid = false
		}
	}

	return
}

func (s *Set) PrintError(ruleName string) {
	if message, ok := s.Messages[ruleName]; ok {
		println(" ✗ " + message)
	} else {
		println(" ✗ " + ruleName)
	}
}
