package user

import (
	"fmt"
	utils "tugasbesar/utils"
	"tugasbesar/vardata"
)

var menuInput int = 0

func PanelUser() {
	utils.ClearConsole()
	fmt.Print("====Welcome User====\n")
	fmt.Print("1. Login\n")
	fmt.Print("2. Registrasi\n")
	fmt.Print("99. Keluar\n")
	fmt.Print("====================\n")
	fmt.Print("Select Menu : ")
	fmt.Scanf("%d\n", &menuInput)
	switch menuInput {
	case 1:
		LoginUser()

	case 2:
		RegisterUser()

	case 99:
		fmt.Print("Exit")

	default:
		PanelUser()
	}
}

func LoginUser() {
	var noTelp, password string
	utils.ClearConsole()
	fmt.Print("=====Login=====\n")
	fmt.Print("No Telp : ")
	fmt.Scanf("%s\n", &noTelp)
	fmt.Print("Password : ")
	fmt.Scanf("%s\n", &password)
	fmt.Print("===============\n")
	userIndex := getUserIndexByTelp(noTelp)
	var isSame = false
	var verif = false
	if userIndex != -1 {
		verif = vardata.UserData[userIndex].IsActive
		isSame = vardata.UserData[userIndex].Password == password
	}

	if verif && isSame {
		//login success
		fmt.Println("SUCCESS")
	} else {
		if !verif {
			fmt.Print("=> Akun anda belum terverifikasi oleh ADMIN\n")
		} else if !isSame {
			fmt.Print("=> No telp atau password anda salah\n")
		} else {
			fmt.Print("=> Akun tidak ditemukan\n")
		}
		fmt.Print("=> Tekan enter untuk kembali\n")
		empty := ""
		fmt.Scanf("%s\n", &empty)
		PanelUser()
	}
}

func getUserIndexByTelp(telp string) int {
	for i := 0; i < len(vardata.UserData); i++ {
		if vardata.UserData[i].NoTelp == telp {
			return i
		}
	}
	return -1
}
