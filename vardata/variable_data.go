package vardata

import (
	"tugasbesar/utils"
)

const CONST_maxUserData int = 1000

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

var UserData [CONST_maxUserData]User

func InitializeDataVariable() {
	UserData = utils.LoadData[[CONST_maxUserData]User]("UserData.json")
	utils.SaveData(UserData, "UserData.json")
}

func AddNewUserData(data User) {
	val := getEmptyUserDataIndex()
	if val != -1 {
		data.UID = val + 1
		UserData[val] = data
	}
	utils.SaveData(UserData, "UserData.json")
}

func getEmptyUserDataIndex() int {
	for i := 0; i < len(UserData); i++ {
		if UserData[i].UID == 0 {
			return i
		}
	}
	return -1
}
