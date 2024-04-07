package main

import "fmt"

func main() {
	fmt.Println("Ты меня все еще любишь?")
	var command string

	_, _ = fmt.Scan(&command)
	if command == "да" {
		fmt.Println("Я тебя тоже люблю... Я очень сожалею о том, что произошло, и хочу с тобой поговорить лично.")
	} else if command == "нет" {
		fmt.Println("Я тебя поняла. Прошу, позволь мне сказать слова благодарности за все то, что было между нами")
	}
}
