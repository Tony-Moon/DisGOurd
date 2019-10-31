package home

import "fmt"

/*
CreateUser doc
*/
func CreateUser() int {
	var nav1, i int

	fmt.Println("Navigated to CreateUser")
	fmt.Println("1 - Go Home")
	fmt.Println("2 - Submit")
	fmt.Scanln(&nav1)

	for i = 0; i >= 1; i = i + 0 {
		switch nav1 {
		case 0:
			fmt.Println("1 - Go Home")
			fmt.Println("2 - Submit")
			fmt.Scanln(&nav1)
		case 1:
			nav1 = 0
			i++
		case 2:
			nav1 = Home()
		default:
			fmt.Println("Try Again")
		}
	}
	return nav1
}
