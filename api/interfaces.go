package api

type DepositStrategy interface {
	Deposit(accountID string, currentBalance int, depositAmount int) error
}
