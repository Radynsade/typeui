package forms

import "strconv"

// Check if string can be converted to any number type except complex.
func isNumber(value string) bool {
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

func maximum(
	value string,
	max float64,
	included bool,
) bool {
	length := valueLength(value)

	if included {
		return length <= max
	}

	return length < max
}

func minimum(
	value string,
	min float64,
	included bool,
) bool {
	length := valueLength(value)

	if included {
		return length >= min
	}

	return length > min
}
