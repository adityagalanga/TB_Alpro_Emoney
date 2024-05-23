package main

import (
	"fmt"
	admin "tugasbesar/module/admin"
	user "tugasbesar/module/user"
	utils "tugasbesar/utils"
)

func main() {
	//start menu
	SplashMenu()
}

func SplashMenu() {
	var inputData int = 0
	utils.ClearConsole()
	FrontMenu()
	fmt.Scan(&inputData)
	if inputData != 99 {
		switch inputData {
		case 1:
			admin.LoginAdmin()

		case 2:
			user.LoginUser()

		default:
			SplashMenu()
		}
	}
}
func FrontMenu() {
	fmt.Print("=====E MONEY=====\n")
	fmt.Print("1. Login as ADMIN\n")
	fmt.Print("2. Login as User\n")
	fmt.Print("99. Quit\n")
	fmt.Print("=================\n")
	fmt.Print("Select Menu : ")
}
