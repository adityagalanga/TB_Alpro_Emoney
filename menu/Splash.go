package menu

import (
	"fmt"
	admin "tugasbesar/admin"
	user "tugasbesar/user"
	utils "tugasbesar/utils"
	vardata "tugasbesar/vardata"
)

func StartApps() {
	//start menu
	vardata.InitializeDataVariable()
	SplashMenu()
}

func SplashMenu() {
	var inputData int = 0
	utils.ClearConsole()
	FrontMenu()
	fmt.Scanf("%d\n", &inputData)
	if inputData != 99 {
		switch inputData {
		case 1:
			admin.PanelAdmin()
			break

		case 2:
			user.PanelUser()
			break

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
