package account

import (
	"movie-api/business/account"

	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Username string
	Password string
	Email    string
}

func FromAccountCore(datafrom account.AccountCore) Account {
	return Account{
		Username: dataFrom.Username,
		Password: dataFrom.Password,
		Email:    dataFrom.Email,
	}
}

func ToAccountCore(datato Account) account.AccountCore {
	return account.AccountCore{
		Username: dataTo.Username,
		Password: dataTo.Password,
		Email:    dataTo.Email,
	}
}
