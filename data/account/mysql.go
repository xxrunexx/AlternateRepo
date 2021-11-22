package account

import (
	"movie-api/business/account"

	"gorm.io/gorm"
)

type AccountData struct {
	db *gorm.DB
}

func NewMySqlAccount(db *gorm.DB) account.Data {
	return &AccountData{
		db: db,
	}
}

func (accData *AccountData) InsertAccount(account *account.AccountCore) error {
	convData := FromAccountCore(*account)

	if err := accData.db.Create(&convData).Error; err != nil {
		return err
	}
	return nil
}
