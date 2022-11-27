package main

import (
	"fmt"

	"github.com/Radynsade/typeui"
	"github.com/Radynsade/typeui/input"
	"github.com/Radynsade/typeui/input/rules"
)

var (
	ageInput   = buildAgeInput()
	emailInput = buildEmailInput()
)

func main() {
	list := typeui.NewList(
		typeui.ListTypeNumbered,
		false,
		"Omniva пакомат (+ €3.99)",
		"Доставка курьером (+ €6.99)",
		"Получить в магазине",
	)

	list.Print()

	email := emailInput.PrintAndRead()
	age := ageInput.PrintAndRead()

	fmt.Printf("Ваш email: %s\n", email)
	fmt.Printf("Ваш возраст: %d\n", int(age))
}

func buildAgeInput() *input.Number {
	field := input.NewNumber(
		"",
		"Введите свой возраст",
		false,
	)

	field.SetMax(120, true)
	field.SetMin(0, true)
	field.SetMessage("isFloat", "Невозможно преобразовать в число.")
	field.SetMessage("maxNumber", "Максимальная продолжительность жизни человека 120 лет.")
	field.SetMessage("minNumber", "Число меньше нуля.")

	return field
}

func buildEmailInput() *input.Text {
	field := input.NewText(
		"",
		"Введите ваш email",
		false,
	)

	field.AddRule(rules.Email())
	field.SetMessage("email", "Неправильная электронная почта.")

	return field
}
