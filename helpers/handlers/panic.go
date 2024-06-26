package handlers

import "fmt"

func CatchPanicHandler() {
	if r := recover(); r != nil {
		fmt.Println("Error occured")
		fmt.Println(r)
	} else {
		fmt.Println("Application running perfectly")
	}
}
