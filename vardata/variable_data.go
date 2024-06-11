package vardata

import (
	"fmt"
	"tugasbesar/utils"
)

const CONST_maxUserData int = 1000
const CONST_maxTransactionData int = 1000
const CONST_maxPayment int = 100
const CONST_maxPaymentItem int = 5

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

type Payment struct {
	UID         int
	PaymentID   int //4 nomor didepan secara DESC
	PaymentName string
	Item        [CONST_maxPaymentItem]PaymentItem
}

type PaymentItem struct {
	ItemName string
	Price    int
}

var UserData [CONST_maxUserData]User
var TransactionData [CONST_maxTransactionData]Transaction
var PaymentData [CONST_maxPayment]Payment

func InitializeDataVariable() {
	UserData = utils.LoadData[[CONST_maxUserData]User]("UserData.json")
	utils.SaveData(UserData, "UserData.json")

	TransactionData = utils.LoadData[[CONST_maxTransactionData]Transaction]("TransactionData.json")
	utils.SaveData(TransactionData, "TransactionData.json")

	PaymentData = utils.LoadData[[CONST_maxPayment]Payment]("PaymentData.json")
	utils.SaveData(PaymentData, "PaymentData.json")
	SortedDESCPaymentIDByPaymentID()
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
		if data.Transaction_type == 1 {
			UserData[GetUserIndexByTelp(data.Transfer_account_source)].Saldo -= data.Nominal
			UserData[GetUserIndexByTelp(data.Transfer_account_target)].Saldo += data.Nominal
		} else if data.Transaction_type == 2 {
			UserData[GetUserIndexByTelp(data.Transfer_account_source)].Saldo -= data.Nominal
		}
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

func SortedDESCPaymentIDByPaymentID() {
	for x := 0; x < CONST_maxPayment; x++ {
		idx := x
		for y := x; y < CONST_maxPayment; y++ {
			if PaymentData[y].PaymentID > PaymentData[idx].PaymentID {
				idx = y
			}
		}
		temp := PaymentData[x]
		PaymentData[x] = PaymentData[idx]
		PaymentData[idx] = temp
	}
}

func GetIndexByPaymentID(valueData int) int {
	minValue := 0
	maxValue := CONST_maxPayment - 1

	for minValue <= maxValue {
		mid := (minValue + maxValue) / 2
		if PaymentData[mid].PaymentID == valueData {
			return mid
		}
		fmt.Println(minValue, " = ", maxValue)
		if PaymentData[mid].PaymentID < valueData {
			maxValue = mid - 1
		} else {
			minValue = mid + 1
		}
	}
	return -1
}
