package main

import (
	"fmt"
)

type Human struct {
	FirstName string
	SeName    string
	Age       int
	Gender    string
	MaxSpeed  int
}

func (h Human) HelloCommand() {
	fmt.Printf("Привет, меня зовут %s %s.\nМне %d лет. Мой пол: %s\n", h.FirstName, h.SeName, h.Age, h.Gender)
}

// Метод, который проверяет совершеннолетие человека
func (h Human) IsAdult() bool {
	return h.Age >= 18
}

type Action struct {
	Human
	ActionType string
}

// Метод, который проверяет, можно ли выполнять действие "Бег"
func (a Action) CanRun() {
	if a.MaxSpeed <= 6 && a.ActionType == "Бег" {
		fmt.Printf("Моя максимальная скорость сликшом мала, я не могу бегать!\n")
	}
	fmt.Printf("Я могу развивать скорость до %d км/ч\n", a.MaxSpeed)
}

func main() {
	a1 := Action{
		Human: Human{
			FirstName: "Николай",
			SeName:    "Петров",
			Age:       16,
			Gender:    "Мужской",
			MaxSpeed:  23,
		},
		ActionType: "Бег",
	}
	a2 := Action{
		Human: Human{
			FirstName: "Екатерина",
			SeName:    "Петрова",
			Age:       28,
			Gender:    "Женский",
			MaxSpeed:  5,
		},
		ActionType: "Бег",
	}

	// Используем методы
	a1.HelloCommand()
	fmt.Println(a1.IsAdult())
	fmt.Println(a1.ActionType)
	a1.CanRun()

	a2.HelloCommand()
	fmt.Println(a2.IsAdult())
	fmt.Println(a2.ActionType)
	a2.CanRun()
}
