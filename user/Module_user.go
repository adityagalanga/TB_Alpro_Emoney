package user

import (
	"fmt"
	"time"
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
	userIndex := vardata.GetUserIndexByTelp(noTelp)
	var isSame = false
	var verif = false
	if userIndex != -1 {
		verif = vardata.UserData[userIndex].IsActive
		isSame = vardata.UserData[userIndex].Password == password
	}

	if verif && isSame {
		//login success
		MainMenuUser(noTelp)
	} else {
		if !verif {
			fmt.Print("=> Akun anda belum terverifikasi oleh ADMIN\n")
		} else if !isSame {
			fmt.Print("=> No telp atau password anda salah\n")
		} else {
			fmt.Print("=> Akun tidak ditemukan\n")
		}
		inputTemplateforBack("Tekan enter untuk kembali")
		PanelUser()
	}
}

func MainMenuUser(telp string) {
	index := vardata.GetUserIndexByTelp(telp)
	utils.ClearConsole()
	fmt.Print("=============================\n")
	fmt.Printf("=> Selamat datang %s\n", vardata.UserData[index].Nama)
	fmt.Printf("=> No Telp %s \n", vardata.UserData[index].NoTelp)
	fmt.Printf("=> Saldo Rp %d\n", vardata.UserData[index].Saldo)
	fmt.Print("=============================\n")
	fmt.Printf("=> 1. Riwayat Transaksi\n")
	fmt.Printf("=> 2. Transfer\n")
	fmt.Printf("=> 3. Pembayaran\n")
	fmt.Printf("=> 4. Keluar\n")
	fmt.Print("=============================\n")
	fmt.Print("Select Menu : ")
	fmt.Scanf("%d\n", &menuInput)
	switch menuInput {
	case 1:
		fmt.Println("Riwayat Transaksi")

	case 2:
		TransferMenu(telp)

	case 3:
		fmt.Println("Pembayaran")

	case 4:
		PanelUser()

	default:
		MainMenuUser(vardata.UserData[index].NoTelp)
	}
}

func TransferMenu(telp string) {
	index := vardata.GetUserIndexByTelp(telp)
	utils.ClearConsole()
	fmt.Print("========Transfer Uang========\n")
	fmt.Print("=============================\n")
	var targetTelp string
	fmt.Print("=> No Telp Tujuan : ")
	fmt.Scanf("%s\n", &targetTelp)
	fmt.Print("===============\n")
	targetIndex := vardata.GetUserIndexByTelp(targetTelp)

	//target not found
	if targetIndex == -1 {
		inputTemplateforBack("No telp tidak ditemukan, Enter untuk kembali ke menu")
		MainMenuUser(telp)
	}
	if telp == targetTelp {
		inputTemplateforBack("Tidak dapat transfer ke akun yang sama, Enter untuk kembali ke menu")
		MainMenuUser(telp)
	}

	// success target
	utils.ClearConsole()
	fmt.Print("========Transfer Uang========\n")
	fmt.Printf("=> No Telp Akun Tujuan : %s \n", vardata.UserData[targetIndex].NoTelp)
	fmt.Print("=============================\n")
	fmt.Printf("=> Info Saldo Sekarang : Rp %d\n", vardata.UserData[index].Saldo)

	var transferValue int = 0
	fmt.Print("Nominal Transfer Rp ")
	fmt.Scanf("%d\n", &transferValue)

	if transferValue == 0 {
		inputTemplateforBack("Tidak dapat transfer Rp 0, Enter untuk kembali ke menu")
		MainMenuUser(telp)
	} else if transferValue > vardata.UserData[index].Saldo {
		inputTemplateforBack("Nominal Transfer melebihi saldo, Enter untuk kembali ke menu")
		MainMenuUser(telp)
	} else {
		ReConfirmTransferMoney(index, targetIndex, transferValue)
	}
}

func ReConfirmTransferMoney(index, targetIndex, transferValue int) {
	utils.ClearConsole()
	fmt.Print("========Transfer Uang========\n")
	fmt.Printf("=> Nama Akun Tujuan : %s \n", vardata.UserData[targetIndex].Nama)
	fmt.Printf("=> No Telp Akun Tujuan : %s \n", vardata.UserData[targetIndex].NoTelp)
	fmt.Printf("=> Total Transfer : Rp %d\n", transferValue)
	fmt.Print("=============================\n")
	var USERPIN int = 0
	fmt.Print("Masukan PIN untuk konfirmasi (0 untuk keluar) : ")
	fmt.Scanf("%d\n", &USERPIN)
	if USERPIN == 0 {
		MainMenuUser(vardata.UserData[index].NoTelp)
	} else {
		if vardata.UserData[index].PIN == USERPIN {
			var data vardata.Transaction
			data.Transaction_type = 1
			data.Transfer_account_source = vardata.UserData[index].NoTelp
			data.Transfer_account_target = vardata.UserData[targetIndex].NoTelp
			data.Nominal = transferValue
			data.Datetime = time.Now().String()

			vardata.AddNewTransferData(data)
			inputTemplateforBack("TRANSFER SUCCESS, enter untuk kembali ke menu")
			MainMenuUser(vardata.UserData[index].NoTelp)
		} else {
			inputTemplateforBack("PIN SALAH, kembali ke main menu")
			MainMenuUser(vardata.UserData[index].NoTelp)
		}
	}
}

func inputTemplateforBack(message string) {
	fmt.Printf("=> %s\n", message)
	empty := ""
	fmt.Scanf("%s\n", &empty)
}
