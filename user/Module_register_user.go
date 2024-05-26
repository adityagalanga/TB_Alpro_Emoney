package user

import (
	"bufio"
	"fmt"
	"os"
	utils "tugasbesar/utils"
	vardata "tugasbesar/vardata"
)

/*
Data User :
    UID      int
	UserID   string
	Password string
	PIN      int
	Nama     string
	NoTelp   int
	Saldo    int
	IsActive bool
*/

func RegisterUser() {
	var inputUser vardata.User
	var email, notelp bool = false, false
	var infoIsDone bool = false
	for !infoIsDone {
		utils.ClearConsole()

		//check error

		headerRegister()
		if email {
			fmt.Print("=> Email Telah Digunakan\n")
		}
		if notelp {
			fmt.Print("=> No Telp telah digunakan\n")
		}
		if email || notelp {
			fmt.Print("\n")
		}

		//input
		fmt.Print("Nama : ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		inputUser.Nama = scanner.Text()

		fmt.Print("No Telepon : ")
		fmt.Scanf("%s\n", &inputUser.NoTelp)

		fmt.Print("No Email : ")
		fmt.Scanf("%s\n", &inputUser.Email)
		fmt.Print("\n==========================\n")
		email = isUserEmailAlreadyUsed(inputUser.Email)
		notelp = isUserNoTelpAlreadyUsed(inputUser.NoTelp)
		if !email && !notelp {
			infoIsDone = true
		}
	}
	utils.ClearConsole()
	previewRegisterUser(inputUser)

}

func previewRegisterUser(data vardata.User) {
	headerRegister()
	fmt.Print("Nama : ", data.Nama)
	fmt.Print("\nNo Telepon : ", data.NoTelp)
	fmt.Print("\nNo Email : ", data.Email)
	fmt.Print("\n==========================\n")
}

func headerRegister() {
	fmt.Print("==========================\n")
	fmt.Print("=====REGISTER E-MONEY=====\n")
	fmt.Print("==========================\n\n")
}

func isUserEmailAlreadyUsed(email string) bool {
	var loop bool = true
	var x int = 0
	for loop {
		if vardata.UserData[x].UID != 0 {
			if vardata.UserData[x].Email == email {
				return true
			}
			x++
		} else {
			loop = false
		}
	}
	return false
}

func isUserNoTelpAlreadyUsed(notelp string) bool {
	var loop bool = true
	var x int = 0
	for loop {
		if vardata.UserData[x].UID != 0 {
			if vardata.UserData[x].NoTelp == notelp {
				return true
			}
			x++
		} else {
			loop = false
		}
	}
	return false
}
