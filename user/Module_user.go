package user

import (
	"fmt"
	"strconv"
	"time"
	utils "tugasbesar/utils"
	"tugasbesar/vardata"
)

func PanelUser() {
	var menuInput int = 0
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
	var menuInput int = 0
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
		ShowTransactionUser(telp, false)

	case 2:
		TransferMenu(telp)

	case 3:
		ShowPaymentUser(telp)

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
			currentTime := time.Now()
			data.Datetime = currentTime.Format("2006-01-02 15:04:05")

			vardata.AddNewTransferData(data)
			inputTemplateforBack("TRANSFER SUCCESS, enter untuk kembali ke menu")
			MainMenuUser(vardata.UserData[index].NoTelp)
		} else {
			inputTemplateforBack("PIN SALAH, kembali ke main menu")
			MainMenuUser(vardata.UserData[index].NoTelp)
		}
	}
}

func ShowTransactionUser(telp string, isDescending bool) {
	var check int = 0
	utils.ClearConsole()
	index := vardata.GetUserIndexByTelp(telp)
	fmt.Print("======Riwayat Transaksi======\n")
	fmt.Print("=============================\n")
	x := 0

	var onCheck bool = true
	var currentData [100]vardata.Transaction
	var N int = 0
	for onCheck {
		if vardata.TransactionData[x].UID != 0 {
			if vardata.TransactionData[x].Transfer_account_source == telp || vardata.TransactionData[x].Transfer_account_target == telp {
				currentData[N] = vardata.TransactionData[x]
				N++
			}
		} else {
			onCheck = false
		}
		x++
	}
	//sort descendingTanggal
	if isDescending {
		for x := 0; x < N; x++ {
			y := x
			for y > 0 {
				if currentData[y-1].Datetime > currentData[y].Datetime {
					temp := currentData[y-1]
					currentData[y-1] = currentData[y]
					currentData[y] = temp
				}
				y--
			}
		}
	} else {
		for x := 0; x < N; x++ {
			idx := x
			for y := x; y < N; y++ {
				if currentData[y].Datetime > currentData[idx].Datetime {
					idx = y
				}
			}
			temp := currentData[x]
			currentData[x] = currentData[idx]
			currentData[idx] = temp
		}
	}

	for x := 0; x < N; x++ {
		if currentData[x].Transaction_type == 1 {
			parsedTime, err := time.Parse("2006-01-02 15:04:05", currentData[x].Datetime)
			if err != nil {
				println(err)
			}
			date := parsedTime.Format("2006-01-02")
			time := parsedTime.Format("15:04:05")
			if currentData[x].Transfer_account_source == telp {
				fmt.Printf("====> Type : Transfer saldo \n")
				fmt.Printf("=> Tanggal : %s %s \n", date, time)
				target := currentData[x].Transfer_account_target
				fmt.Printf("=> Nama Penerima : %s \n", vardata.UserData[vardata.GetUserIndexByTelp(target)].Nama)
				fmt.Printf("=> No Telp Penerima : %s \n", currentData[x].Transfer_account_target)
				fmt.Printf("=> Total Transfer : Rp %d\n", currentData[x].Nominal)
				fmt.Print("=============================\n")
			} else {
				fmt.Printf("====> Type : Menerima saldo \n")
				fmt.Printf("=> Tanggal : %s %s \n", date, time)
				target := currentData[x].Transfer_account_source
				fmt.Printf("=> Nama Pengirim : %s \n", vardata.UserData[vardata.GetUserIndexByTelp(target)].Nama)
				fmt.Printf("=> No Telp Pengirim : %s \n", currentData[x].Transfer_account_source)
				fmt.Printf("=> Total Transfer : Rp %d\n", currentData[x].Nominal)
				fmt.Print("=============================\n")
			}
		} else if currentData[x].Transaction_type == 2 {
			parsedTime, err := time.Parse("2006-01-02 15:04:05", currentData[x].Datetime)
			if err != nil {
				println(err)
			}
			date := parsedTime.Format("2006-01-02")
			time := parsedTime.Format("15:04:05")
			fmt.Printf("====> Type : Payment Virtual Account \n")
			fmt.Printf("=> Tanggal : %s %s \n", date, time)
			fmt.Printf("=> Nomor Payment : %s \n", currentData[x].Payment_content)
			fmt.Printf("=> Nama Payment : %s \n", currentData[x].Transfer_account_target)
			fmt.Printf("=> Item Payment : %s \n", currentData[x].Payment_id)
			fmt.Printf("=> Total Payment : Rp %d\n", currentData[x].Nominal)
			fmt.Print("=============================\n")
		}
	}
	fmt.Printf("=> 1. Sort ASC\n")
	fmt.Printf("=> 2. Sort DESC\n")
	fmt.Printf("=> 3. Exit\n")
	fmt.Print("Select Menu : ")
	fmt.Scanf("%d\n", &check)
	switch check {
	case 1:
		ShowTransactionUser(telp, true)

	case 2:
		ShowTransactionUser(telp, false)

	case 3:
		MainMenuUser(vardata.UserData[index].NoTelp)

	default:
		MainMenuUser(vardata.UserData[index].NoTelp)
	}
}

func ShowPaymentUser(telp string) {
	index := vardata.GetUserIndexByTelp(telp)
	utils.ClearConsole()
	fmt.Print("========PaymentEmoney========\n")
	fmt.Print("=============================\n")
	var targetVA string
	fmt.Print("=> No Virtual Account : ")
	fmt.Scanf("%s\n", &targetVA)
	fmt.Print("===============\n")
	firstFourDigits, err := strconv.Atoi(targetVA[:4])
	if err != nil {
		fmt.Println("Error:", err)
	}
	targetIndex := vardata.GetIndexByPaymentID(firstFourDigits)
	//target not found
	if targetIndex == -1 {
		inputTemplateforBack("No PaymentID tidak ditemukan, Enter untuk kembali ke menu")
		MainMenuUser(telp)
	}
	// if telp == targetVA {
	// 	inputTemplateforBack("Tidak dapat transfer ke akun yang sama, Enter untuk kembali ke menu")
	// 	MainMenuUser(telp)
	// }

	// success target
	utils.ClearConsole()
	fmt.Print("========PaymentEmoney========\n")
	fmt.Printf("=> No Virtual Account : %s \n", targetVA)
	fmt.Printf("=> Nama Payment : %s \n", vardata.PaymentData[targetIndex].PaymentName)
	fmt.Print("=============================\n")
	for x := 0; x < vardata.CONST_maxPaymentItem; x++ {
		if vardata.PaymentData[targetIndex].Item[x].ItemName != "" {
			fmt.Printf("=> %d. %s = Rp %d\n", x+1, vardata.PaymentData[targetIndex].Item[x].ItemName, vardata.PaymentData[targetIndex].Item[x].Price)
		}
	}
	fmt.Print("=============================\n")
	var itemCheck int
	fmt.Print("Select Menu : ")
	fmt.Scanf("%d\n", &itemCheck)
	if itemCheck-1 <= vardata.CONST_maxPaymentItem && itemCheck >= 0 {
		if vardata.PaymentData[targetIndex].Item[itemCheck-1].ItemName != "" {
			if vardata.PaymentData[targetIndex].Item[itemCheck-1].Price > vardata.UserData[index].Saldo {
				inputTemplateforBack("Nominal Payment melebihi saldo, Enter untuk kembali ke menu")
				MainMenuUser(telp)
			} else {
				ReConfirmPayment(targetVA, index, targetIndex, itemCheck-1)
			}
		} else {
			MainMenuUser(telp)
		}
	} else {
		MainMenuUser(telp)
	}
}
func ReConfirmPayment(targetVA string, index, targetIndex, itemCheck int) {
	utils.ClearConsole()
	fmt.Print("========PaymentEmoney========\n")
	fmt.Printf("=> No Virtual Account : %s \n", targetVA)
	fmt.Printf("=> Nama Payment : %s \n", vardata.PaymentData[targetIndex].PaymentName)
	fmt.Printf("=> Nama Item : %s \n", vardata.PaymentData[targetIndex].Item[itemCheck].ItemName)
	fmt.Printf("=> Total Pembayaran : Rp %d\n", vardata.PaymentData[targetIndex].Item[itemCheck].Price)
	fmt.Print("=============================\n")
	var USERPIN int = 0
	fmt.Print("Masukan PIN untuk konfirmasi (0 untuk keluar) : ")
	fmt.Scanf("%d\n", &USERPIN)
	if USERPIN == 0 {
		MainMenuUser(vardata.UserData[index].NoTelp)
	} else {
		if vardata.UserData[index].PIN == USERPIN {
			var data vardata.Transaction
			data.Transaction_type = 2
			data.Transfer_account_source = vardata.UserData[index].NoTelp
			data.Transfer_account_target = vardata.PaymentData[targetIndex].PaymentName
			data.Nominal = vardata.PaymentData[targetIndex].Item[itemCheck].Price
			data.Payment_id = vardata.PaymentData[targetIndex].Item[itemCheck].ItemName
			data.Payment_content = targetVA
			currentTime := time.Now()
			data.Datetime = currentTime.Format("2006-01-02 15:04:05")
			vardata.AddNewTransferData(data)
			inputTemplateforBack("PAYMENT SUCCESS, enter untuk kembali ke menu")
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
