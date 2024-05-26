package user

import (
	"fmt"
	utils "tugasbesar/utils"
)

var menuInput int = 0

func LoginUser() {
	utils.ClearConsole()
	fmt.Print("=====User Login=====\n")
	fmt.Print("1. Login\n")
	fmt.Print("2. Registrasi\n")
	fmt.Print("99. Keluar\n")
	fmt.Print("====================\n")
	fmt.Print("Select Menu : ")
	fmt.Scanf("%d\n", &menuInput)
	switch menuInput {
	case 1:
		fmt.Print("Login")

	case 2:
		RegisterUser()

	case 99:
		fmt.Print("Exit")

	default:
		LoginUser()
	}
}
