package dto

type DepositMoneyRequest struct {
	WalletID int `json:"wallet_id"`
	Amount   int `json:"amount"`
}
