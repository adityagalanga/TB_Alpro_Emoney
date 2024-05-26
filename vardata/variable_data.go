package vardata

import (
	utils "tugasbesar/utils"
)

const CONST_maxUserData int = 1000

type User struct {
	UID      int
	UserID   string
	Password string
	PIN      int
	Nama     string
	NoTelp   int
	Saldo    int
	IsActive bool
}

var UserData [CONST_maxUserData]User

func InitializeDataVariable() {
	UserData = utils.LoadData[[CONST_maxUserData]User]("UserData.json")
}
