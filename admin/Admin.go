package admin

import "fmt"

/*
Admin doc
*/
func Admin() int {
	var nav1, i int

	fmt.Println("Navigated to Admin")
	fmt.Println("1 - Logout")
	fmt.Scanln(&nav1)

	for i = 0; i >= 1; i = i + 0 {
		switch nav1 {
		case 0:
			fmt.Println("1 - Logout")
			fmt.Scanln(&nav1)
		case 1:
			nav1 = 0
			i++
		default:
			fmt.Println("Try Again")
		}
	}
	return nav1
}
