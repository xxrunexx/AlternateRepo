package account

type AccountBusiness struct {
	accountData Data
}

func NewBusinessAccount(accountData Data) Business {
	return &AccountBusiness{
		accountData: accountData,
	}
}

func (accBusiness *AccountBusiness) CreateAccount(accData AccountCore) error {
	// if err := accBusiness.accountData.InsertAccount(accData); err != nil {
	if err := accBusiness.accountData.InsertAccount(&accData); err != nil {
		return err
	}
	return nil
}
