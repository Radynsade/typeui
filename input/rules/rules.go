package rules

import (
	"net/url"
	"reflect"
	"regexp"
	"strconv"
)

var (
	emailRegexp, _ = regexp.Compile("^(([^<>()[\\]\\.,;:\\s@\"]+(\\.[^<>()[\\]\\.,;:\\s@\"]+)*)|(\".+\"))@((\\[[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\])|(([a-zA-Z\\-0-9]+\\.)+[a-zA-Z]{2,}))$")
)

// First return value is the name of the rule, second is result (valid or not).
type Rule func(value string) (string, bool)

func IsInt() Rule {
	const ruleName = "isInt"

	return func(value string) (string, bool) {
		return ruleName, isInteger(value)
	}
}

func IsFloat() Rule {
	const ruleName = "isFloat"

	return func(value string) (string, bool) {
		return ruleName, isFloat(value)
	}
}

func Email() Rule {
	const ruleName = "email"

	return func(value string) (string, bool) {
		return ruleName, emailRegexp.MatchString(value)
	}
}

func Required() Rule {
	const ruleName = "required"

	return func(value string) (string, bool) {
		return ruleName, !reflect.ValueOf(value).IsZero()
	}
}

func Url() Rule {
	const ruleName = "url"

	return func(value string) (string, bool) {
		_, err := url.ParseRequestURI(value)

		return ruleName, err == nil
	}
}

func Regexp(regex *regexp.Regexp) Rule {
	const ruleName = "regexp"

	return func(value string) (string, bool) {
		return ruleName, regex.MatchString(value)
	}
}

func MaxNumber(
	max float64,
	included bool,
) Rule {
	const ruleName = "maxNumber"

	return func(value string) (string, bool) {
		num, err := strconv.ParseFloat(value, 64)

		if err != nil {
			return ruleName, false
		}

		return ruleName, maxNumber(num, max, included)
	}
}

func MinNumber(
	min float64,
	included bool,
) Rule {
	const ruleName = "minNumber"

	return func(value string) (string, bool) {
		num, err := strconv.ParseFloat(value, 64)

		if err != nil {
			return ruleName, false
		}

		return ruleName, minNumber(num, min, included)
	}
}

func MaxLength(
	max int,
	included bool,
) Rule {
	const ruleName = "maxLength"

	return func(value string) (string, bool) {
		return ruleName, maxLength[string](value, max, included)
	}
}

func MinLength(
	min int,
	included bool,
) Rule {
	const ruleName = "minLength"

	return func(value string) (string, bool) {
		return ruleName, minLength[string](value, min, included)
	}
}

func OneOf(values ...string) Rule {
	const ruleName = "oneOf"

	return func(value string) (string, bool) {
		for _, v := range values {
			if v == value {
				return ruleName, true
			}
		}

		return ruleName, false
	}
}
