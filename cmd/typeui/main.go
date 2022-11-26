package main

import (
	"github.com/Radynsade/typeui"
	"github.com/Radynsade/typeui/forms"
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

	passField := passwordField()
	passField.PrintAndRead()

	println(passField.Field.Value)
}

func passwordField() *typeui.BaseInput {
	field := forms.NewField(
		"password",
		forms.Required(),
		forms.Min(5, true),
	)

	return typeui.NewBaseInput(
		"Пароль",
		"Введите новый пароль",
		field,
		true,
	)
}
