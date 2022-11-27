package rules

import "strconv"

// Check if string can be converted to any number type except complex.
func isFloat(value string) bool {
	_, err := strconv.ParseFloat(value, 64)

	return err == nil
}

func isInteger(value string) bool {
	_, err := strconv.Atoi(value)

	return err == nil
}

func valueLength(value string) float64 {
	floatnum, err := strconv.ParseFloat(value, 64)

	if err == nil {
		return floatnum
	}

	return float64(len(value))
}

func maxNumber(
	value float64,
	max float64,
	included bool,
) bool {
	if included {
		return value <= max
	}

	return value < max
}

func minNumber(
	value float64,
	min float64,
	included bool,
) bool {
	if included {
		return value >= min
	}

	return value > min
}

func maxLength[T comparable](
	value string,
	max int,
	included bool,
) bool {
	length := len(value)

	if included {
		return length <= max
	}

	return length < max
}

func minLength[T comparable](
	value string,
	min int,
	included bool,
) bool {
	length := len(value)

	if included {
		return length >= min
	}

	return length > min
}
