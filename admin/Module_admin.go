package admin

import (
	"fmt"
	"tugasbesar/utils"
	"tugasbesar/vardata"
)

func PanelAdmin() {
	var menuInput int = 0
	utils.ClearConsole()
	fmt.Print("====Welcome Admin====\n")
	fmt.Print("1. Login\n")
	fmt.Print("99. Keluar\n")
	fmt.Print("====================\n")
	fmt.Print("Select Menu : ")
	fmt.Scanf("%d\n", &menuInput)
	switch menuInput {
	case 1:
		LoginAdmin()
		break
	default:
		fmt.Print("Exit")
	}
}

func LoginAdmin() {
	var noTelp, password string
	utils.ClearConsole()
	fmt.Print("=====Login=====\n")
	fmt.Print("No Telp : ")
	fmt.Scanf("%s\n", &noTelp)
	fmt.Print("Password : ")
	fmt.Scanf("%s\n", &password)
	fmt.Print("===============\n")
	userIndex := vardata.GetUserIndexByTelp(noTelp)

	var isMatch, isVerified, isAdmin = false, false, false

	if userIndex != -1 {
		isMatch = vardata.UserData[userIndex].Password == password
		isVerified = vardata.UserData[userIndex].IsActive
		isAdmin = vardata.UserData[userIndex].IsAdmin

		if !isMatch {
			fmt.Print("=> No telp atau password anda salah\n")
			backToMenu("LoginAdmin")
		}

		if !isVerified {
			fmt.Print("=> Akun anda belum terverifikasi oleh ADMIN\n")
			backToMenu("LoginAdmin")
		}

		if !isAdmin {
			fmt.Print("=> Akun tidak memilki has akses sebagai ADMIN\n")
			backToMenu("LoginAdmin")
		}

		MainMenuAdmin()
	} else {
		fmt.Printf("=> Akun tidak ditemukan\n")
		backToMenu("LoginAdmin")
	}
}

func MainMenuAdmin() {
	var menuInput int = 0
	utils.ClearConsole()
	fmt.Print("=============================\n")
	fmt.Printf("=> Selamat datang Admin\n")
	fmt.Print("=============================\n")
	fmt.Printf("=> 1. Lihat Daftar Pengguna\n")
	fmt.Printf("=> 2. Verifikasi Pengguna\n")
	fmt.Printf("=> 99. Keluar\n")
	fmt.Print("=============================\n")
	fmt.Print("Select Menu : ")
	fmt.Scanf("%d\n", &menuInput)
	switch menuInput {
	case 1:
		showUserListMenu()
		break
	case 2:
		verifyUserMenu()
		break
	default:
		backToMenu("PanelAdmin")
	}
}

func showUserListMenu() {
	var userList [100]vardata.User
	var userDataIndex, userListIndex = 0, 0
	var onCheck = true

	for onCheck {
		if vardata.UserData[userDataIndex].UID != 0 {
			if !vardata.UserData[userDataIndex].IsAdmin {
				userList[userListIndex] = vardata.UserData[userDataIndex]
				userListIndex++
			}
		} else {
			onCheck = false
		}

		userDataIndex++
	}

	utils.ClearConsole()
	fmt.Print("=======Daftar Pengguna=======\n")
	fmt.Print("=============================\n")
	for i := 0; i < userListIndex; i++ {
		var status string
		if userList[i].IsActive {
			status = "SUDAH DIVERIFIKASI"
		} else {
			status = "BELUM DIVERIFIKASI"
		}

		fmt.Printf("=> UID Pengguna     : %d \n", userList[i].UID)
		fmt.Printf("=> Nama Pengguna    : %s \n", userList[i].Nama)
		fmt.Printf("=> Email Pengguna   : %s \n", userList[i].Email)
		fmt.Printf("=> No Telp Pengguna : %s \n", userList[i].NoTelp)
		fmt.Printf("=> Status Pengguna  : %s \n", status)
		fmt.Print("=============================\n")
	}
	fmt.Printf("\n%d data telah ditampilkan\n\n", userListIndex)
	backToMenu("MainMenuAdmin")
}

func verifyUserMenu() {
	var noTelp string
	utils.ClearConsole()
	fmt.Print("=====Verifikasi Pengguna=====\n")
	fmt.Print("=============================\n")
	fmt.Print("No Telp Pengguna : ")
	fmt.Scanf("%s\n", &noTelp)

	userIndex := vardata.GetUserIndexByTelp(noTelp)
	if userIndex != -1 {
		userData := vardata.UserData[userIndex]
		if !userData.IsActive {
			var confirm string
			fmt.Print("\n========Data Pengguna========\n")
			fmt.Print("=============================\n")
			fmt.Printf("=> UID Pengguna     : %d \n", userData.UID)
			fmt.Printf("=> Nama Pengguna    : %s \n", userData.Nama)
			fmt.Printf("=> Email Pengguna   : %s \n", userData.Email)
			fmt.Printf("=> No Telp Pengguna : %s \n", userData.NoTelp)
			fmt.Print("=============================\n")
			fmt.Print("(Y/N) untuk konfirmasi: ")
			fmt.Scanf("%s\n", &confirm)
			if confirm == "Y" || confirm == "y" {
				vardata.ActivateUserData(userData)
				fmt.Printf("=> Verifikasi berhasil\n")
				backToMenu("MainMenuAdmin")
			} else {
				fmt.Printf("=> Verifikasi gagal\n")
				backToMenu("MainMenuAdmin")
			}
		} else {
			fmt.Printf("=> Akun telah aktif\n")
			backToMenu("MainMenuAdmin")
		}
	} else {
		fmt.Printf("=> Akun tidak ditemukan\n")
		backToMenu("MainMenuAdmin")
	}
}

func backToMenu(targetMenu string) {
	inputTemplateForBack("Tekan enter untuk kembali")

	switch targetMenu {
	case "PanelAdmin":
		PanelAdmin()
		break
	case "LoginAdmin":
		LoginAdmin()
		break
	case "MainMenuAdmin":
		MainMenuAdmin()
		break
	default:
		PanelAdmin()
	}
}

func inputTemplateForBack(message string) {
	fmt.Printf("=> %s\n", message)
	empty := ""
	fmt.Scanf("%s\n", &empty)
}
