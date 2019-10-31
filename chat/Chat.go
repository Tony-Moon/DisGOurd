package chat

import (
	"fmt"

	"github.com/Tony-Moon/DisGOurd/home"
)

/*
Chat doc
*/
func Chat() {
	var nav1, i int

	fmt.Println("Navigated to Chat")
	fmt.Println("1 - Go Home")
	fmt.Println("2 - Logout")
	fmt.Scanln(&nav1)

	for i = 0; i >= 1; i = i + 0 {
		switch nav1 {
		case 0:
			fmt.Println("1 - Go Home")
			fmt.Println("2 - Logout")
			fmt.Scanln(&nav1)
		case 1:
			nav1 = home.Home()
		case 2:
			nav1 = 0
			i++
		default:
			fmt.Println("Try Again")
		}
	}
	return nav1
}
