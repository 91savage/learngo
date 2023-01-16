package accounts

// Account struct
type Account struct {
	owner string
	balance int
}

// NewAccount creates Account
func NewAccount(owner string) *Account{
	account := Account{owner: owner, balance: 0}
	return &account
	// & : Address
	// * : See trough
}
// Deposit x amount on your account
func (a *Account) Deposit(amount int) {  // (a Account) Go에서 receiver 라고 함.
	a.balance += amount
}

// Balance of your account
func (a *Account) Balance() int {
	return a.balance
}