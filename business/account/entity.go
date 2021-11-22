package account

// Domain layer / entity yang paling luar
type AccountCore struct {
	ID       int
	Username string
	Password string
	Email    string
}

// Data / Repo
type Data interface {
	InsertAccount(account *AccountCore) (*AccountCore, error)
}

// Business / Service
type Business interface {
	CreateAccount(account *AccountCore) (*AccountCore, error)
}
