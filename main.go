package main

import (
	"fmt"

	"github.com/Tony-Moon/DisGOurd/admin"
	"github.com/Tony-Moon/DisGOurd/home"
)

func main() {
	var i, nav1 int
	i = 0

	fmt.Println("Project Started without HTML")
	fmt.Println("1 - User Login")
	fmt.Println("2 - Admin Login")
	fmt.Println("3 - Create Account")
	fmt.Println("4 - Quit")
	fmt.Scanln(&nav1)

	for i = 0; i >= 1; i = i + 0 {
		switch nav1 {
		case 0:
			fmt.Println("1 - User Login")
			fmt.Println("2 - Admin Login")
			fmt.Println("3 - Create Account")
			fmt.Println("4 - Quit")
			fmt.Scanln(&nav1)
		case 1:
			nav1 = home.UserLogin()
		case 2:
			nav1 = admin.ALogin()
		case 3:
			nav1 = home.CreateUser()
		case 4:
			i++
		default:
			fmt.Println("Try Again")
		}
	}
}
