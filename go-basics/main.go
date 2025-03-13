package main

import (
	"fmt"
	"go-basics/foods"
	"go-basics/messages"
	"strings"
)

func main() {

	fmt.Println(strings.ToUpper(DisplayCustomMessage("Sow", "Go's very faster")))

	fmt.Println(messages.DisplayCustomMessage("Venkat", "I love Golang!"))

	order, err := foods.FoodOrder(-3, "Pizzas")

	if err != nil {
		fmt.Println("Veuillez vérifier les quantités demandées", err)
	}

	fmt.Println(order)

}

func DisplayCustomMessage(author, message string) string {

	return fmt.Sprintf("%v : %v", author, message)
}
