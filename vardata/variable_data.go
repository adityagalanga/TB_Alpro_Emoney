package vardata

import (
	"tugasbesar/utils"
)

const CONST_maxUserData int = 1000

type User struct {
	UID      int
	UserID   string
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
	utils.SaveData[[CONST_maxUserData]User](UserData, "UserData.json")
}
