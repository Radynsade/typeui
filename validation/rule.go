package validation

type Rule func(value interface{}) bool

type RuleError struct {
}
