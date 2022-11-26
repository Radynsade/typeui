package forms

import (
	"net/mail"
	"net/url"
	"reflect"
)

// First return value is the name of the rule, second is result (valid or not).
type Validator func(value string) (string, bool)

func IsInt() Validator {
	const validatorName = "isInt"

	return func(value string) (string, bool) {
		return validatorName, isInteger(value)
	}
}

func IsFloat() Validator {
	const validatorName = "isFloat"

	return func(value string) (string, bool) {
		return validatorName, isNumber(value)
	}
}

func Max(
	max float64,
	included bool,
) Validator {
	const validatorName = "max"

	return func(value string) (string, bool) {
		return validatorName, maximum(value, max, included)
	}
}

func Min(
	min float64,
	included bool,
) Validator {
	const validatorName = "min"

	return func(value string) (string, bool) {
		return validatorName, minimum(value, min, included)
	}
}

func Email() Validator {
	const validatorName = "email"

	return func(value string) (string, bool) {
		_, err := mail.ParseAddress(value)

		return validatorName, err == nil
	}
}

func Required() Validator {
	const validatorName = "required"

	return func(value string) (string, bool) {
		return validatorName, !reflect.ValueOf(value).IsZero()
	}
}

func Url() Validator {
	const validatorName = "url"

	return func(value string) (string, bool) {
		_, err := url.ParseRequestURI(value)

		return validatorName, err == nil
	}
}
