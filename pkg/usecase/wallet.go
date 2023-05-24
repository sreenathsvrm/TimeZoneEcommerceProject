package usecase

import (

)

// func (c *Orderusecase) GetUserWallet(ctx context.Context, userID uint) (wallet domain.Wallet, err error) {

// 	// first find the user wallet
// 	wallet, err = c.orderRepo.FindWalletByUserID(ctx, userID)
// 	if err != nil {
// 		return wallet, err
// 	} else if wallet.ID == 0 { // if user have no wallet then create a wallet for user
// 		wallet.ID, err = c.orderRepo.SaveWallet(ctx, userID)
// 		if err != nil {
// 			return wallet, err
// 		}
// 	}

// 	log.Printf("successfully got user wallet with wallet_id %v for user user_id %v", wallet.ID, userID)
// 	return wallet, nil
// }

// func (c *Orderusecase) GetUserWalletTransactions(ctx context.Context,userID uint, pagination urequest.Pagination) (transactions []domain.Transaction, err error) {

// 	// first find the user wallet
// 	wallet, err := c.orderRepo.FindWalletByUserID(ctx, userID)
// 	if err != nil {
// 		return transactions, err
// 	} else if wallet.ID == 0 {
// 		return transactions, fmt.Errorf("there is no wallet for user with user_id %v for showing transaction", userID)
// 	}

// 	// then find the transactions by wallet_id
// 	transactions, err = c.orderRepo.FindWalletTransactions(ctx, wallet.ID, pagination)

// 	if err != nil {
// 		return transactions, err
// 	}

// 	return transactions, nil
// }
