package repository

import (
	"context"
	"ecommerce/pkg/commonhelp/urequest"
	"ecommerce/pkg/domain"
	"fmt"
	"time"
)

func (c *OrderDB) FindWalletByUserID(ctx context.Context, userID uint) (wallet domain.Wallet, err error) {
	query := `SELECT * FROM wallets WHERE user_id = $1`

	err = c.DB.Raw(query, userID).Scan(&wallet).Error

	if err != nil {
		return wallet, fmt.Errorf("faild to find wallet for wallet_id %v", userID)
	}

	return wallet, nil
}

// create a new wallet for user
func (c *OrderDB) SaveWallet(ctx context.Context, userID uint) (walletID uint, err error) {

	var wallet domain.Wallet
	query := `INSERT INTO wallets (user_id,total_amount) VALUES ($1, $2) RETURNING wallet_id`
	err = c.DB.Raw(query, userID, 0).Scan(&wallet).Error

	walletID = wallet.ID

	return walletID, err
}

func (c *OrderDB) UpdateWallet(ctx context.Context, walletID, upateTotalAmount uint) error {

	query := `UPDATE wallets SET total_amount = $1 WHERE wallet_id = $2`
	err := c.DB.Exec(query, upateTotalAmount, walletID).Error

	return err
}

func (c *OrderDB) SaveWalletTransaction(ctx context.Context, walletTrx domain.Transaction) error {

	trxDate := time.Now()
	query := `INSERT INTO transactions (wallet_id, transaction_date, amount, transaction_type) 
	VALUES ($1, $2, $3, $4)`
	err := c.DB.Exec(query, walletTrx.WalletID, trxDate, walletTrx.Amount, walletTrx.TransactionType).Error

	return err
}

// find wallet transaction history

func (c *OrderDB) FindWalletTransactions(ctx context.Context, walletID uint, pagination urequest.Pagination) (transaction []domain.Transaction, err error) {

	limit := pagination.PerPage
	offset := (pagination.Page- 1) * limit

	query := `SELECT * FROM transactions WHERE id = $1
	ORDER BY transaction_date DESC LIMIT $2 OFFSET $3`

	err = c.DB.Raw(query, walletID, limit, offset).Scan(&transaction).Error

	if err != nil {
		return transaction, fmt.Errorf("faild get transactions of this wallet_id %v", walletID)
	}

	return transaction, nil
}
