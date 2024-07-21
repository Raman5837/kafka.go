package utils

import "gorm.io/gorm"

// A Wrapper To Get Atomic Transaction Functionality
func WithTransactionAtomic(DB *gorm.DB, inner func(transaction *gorm.DB) error) error {

	newTransaction := DB.Begin()

	if newTransaction.Error != nil {
		return newTransaction.Error
	}

	// If Something Breaks, Rollback.
	defer func() {
		if __recover := recover(); __recover != nil {
			newTransaction.Rollback()
		}
	}()

	// If Wrapped Function Returns An Error, Rollback.
	if exception := inner(newTransaction); exception != nil {
		newTransaction.Rollback()
	}

	// If Fails At DB Level, Rollback.
	if dbErr := newTransaction.Commit().Error; dbErr != nil {
		newTransaction.Rollback()
	}

	return nil

}
