package vardata

import (
	"tugasbesar/utils"
)

const CONST_maxUserData int = 1000
const CONST_maxTransactionData int = 1000

type User struct {
	UID      int
	Password string
	PIN      int
	Nama     string
	Email    string
	NoTelp   string
	Saldo    int
	IsActive bool
}

type Transaction struct {
	UID                     int
	Transaction_type        int //1 transfer, 2 pembayaran
	Transfer_account_source string
	Transfer_account_target string
	Nominal                 int
	Datetime                string
	Payment_id              string
	Payment_content         string
}

var UserData [CONST_maxUserData]User
var TransactionData [CONST_maxTransactionData]Transaction

func InitializeDataVariable() {
	UserData = utils.LoadData[[CONST_maxUserData]User]("UserData.json")
	utils.SaveData(UserData, "UserData.json")

	TransactionData = utils.LoadData[[CONST_maxTransactionData]Transaction]("TransactionData.json")
	utils.SaveData(TransactionData, "TransactionData.json")
}

func AddNewUserData(data User) {
	val := getEmptyUserDataIndex()
	if val != -1 {
		data.UID = val + 1
		UserData[val] = data
	}
	utils.SaveData(UserData, "UserData.json")
}

func AddNewTransferData(data Transaction) {
	val := getEmptyTransactionIndex()
	if val != -1 {
		data.UID = val + 1
		TransactionData[val] = data
		UserData[GetUserIndexByTelp(data.Transfer_account_source)].Saldo -= data.Nominal
		UserData[GetUserIndexByTelp(data.Transfer_account_target)].Saldo += data.Nominal
	}
	utils.SaveData(UserData, "UserData.json")
	utils.SaveData(TransactionData, "TransactionData.json")
}

func getEmptyTransactionIndex() int {
	for i := 0; i < len(TransactionData); i++ {
		if TransactionData[i].UID == 0 {
			return i
		}
	}
	return -1
}

func getEmptyUserDataIndex() int {
	for i := 0; i < len(UserData); i++ {
		if UserData[i].UID == 0 {
			return i
		}
	}
	return -1
}

func GetUserIndexByTelp(telp string) int {
	for i := 0; i < len(UserData); i++ {
		if UserData[i].NoTelp == telp {
			return i
		}
	}
	return -1
}
