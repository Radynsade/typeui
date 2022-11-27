package rules

func Validate(
	input string,
	rules []Rule,
) (results map[string]bool, allValid bool) {
	results = make(map[string]bool)
	allValid = true

	for _, rule := range rules {
		ruleName, valid := rule(input)
		results[ruleName] = valid

		if !valid {
			allValid = false
		}
	}

	return
}
