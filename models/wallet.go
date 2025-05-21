package models

type WalletRequest struct {
    UserID string `json:"user_id"`
}

type WalletResponse struct {
    UserID  string `json:"user_id"`
    Address string `json:"wallet_address"`
}
