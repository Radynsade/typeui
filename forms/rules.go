package forms

import (
	"net/mail"
	"net/url"
	"reflect"
)

// First return value is the name of the rule, second is result (valid or not).
type Rule func(value string) (string, bool)

func RuleIsInt() Rule {
	const ruleName = "isInt"

	return func(value string) (string, bool) {
		return ruleName, isInteger(value)
	}
}

func RuleIsFloat() Rule {
	const ruleName = "isFloat"

	return func(value string) (string, bool) {
		return ruleName, isNumber(value)
	}
}

func RuleMax(
	max float64,
	included bool,
) Rule {
	const ruleName = "max"

	return func(value string) (string, bool) {
		return ruleName, maximum(value, max, included)
	}
}

func RuleMin(
	min float64,
	included bool,
) Rule {
	const ruleName = "min"

	return func(value string) (string, bool) {
		return ruleName, minimum(value, min, included)
	}
}

func RuleEmail() Rule {
	const ruleName = "email"

	return func(value string) (string, bool) {
		_, err := mail.ParseAddress(value)

		return ruleName, err == nil
	}
}

func RuleRequired() Rule {
	const ruleName = "required"

	return func(value string) (string, bool) {
		return ruleName, !reflect.ValueOf(value).IsZero()
	}
}

func RuleUrl() Rule {
	const ruleName = "url"

	return func(value string) (string, bool) {
		_, err := url.ParseRequestURI(value)

		return ruleName, err == nil
	}
}
