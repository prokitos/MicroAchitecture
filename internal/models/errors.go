package models

import "fmt"

// обработка ошибок
func CheckError(err error) {
	defer panicProcessing()
	if err != nil {
		panic(err)
	}
}

func panicProcessing() {

	if a := recover(); a != nil {
		fmt.Println("ERROR, PANIC", a)
	}
}
