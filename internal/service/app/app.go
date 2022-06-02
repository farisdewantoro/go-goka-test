package app

type Application struct {
	DepositMoney           DepositMoneyHandler
	GetBalanceWalletDetail GetBalanceWalletDetailHandler
}
