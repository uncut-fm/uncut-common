package model

import (
	common_model "github.com/uncut-fm/uncut-common/model"
	"reflect"
	"time"
)

type Wallet struct {
	ID              int
	Name            string
	Description     string
	WalletAddress   string
	Provider        string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	BecamePrimaryAt time.Time
	LastSyncedAt    time.Time

	User         *User
	NFTOwners    []*NFTOwner
	Transactions []*Transaction
}

func NewWalletsListFromCommonWallets(commonWallets []*common_model.Wallet, user User) []*Wallet {
	wallets := make([]*Wallet, len(commonWallets))

	// remove the user's wallets to avoid infinite recursion
	user.Wallets = nil

	for i := range commonWallets {
		wallets[i] = NewWalletFromCommonWallet(commonWallets[i], user)
	}

	return wallets
}

func NewWalletFromCommonWallet(commonWallet *common_model.Wallet, user User) *Wallet {
	return &Wallet{
		ID:            commonWallet.ID,
		Name:          commonWallet.Name,
		Description:   commonWallet.Description,
		WalletAddress: commonWallet.WalletAddress,
		Provider:      commonWallet.Provider,
		CreatedAt:     commonWallet.CreatedAt,
		UpdatedAt:     commonWallet.UpdatedAt,
		LastSyncedAt:  commonWallet.LastSyncedAt,
		User:          &user,
	}
}

// GetPropertiesInMap returns a map of the wallet's properties; keys are in camelCase
func (w *Wallet) GetPropertiesInMap() map[string]interface{} {
	return map[string]interface{}{
		"id":              w.ID,
		"name":            w.Name,
		"description":     w.Description,
		"walletAddress":   w.WalletAddress,
		"provider":        w.Provider,
		"createdAt":       w.CreatedAt.Format("2006-01-02 15:04:05 MST"),
		"updatedAt":       w.UpdatedAt.Format("2006-01-02 15:04:05 MST"),
		"lastSyncedAt":    w.LastSyncedAt.Format("2006-01-02 15:04:05 MST"),
		"becamePrimaryAt": w.BecamePrimaryAt.Format("2006-01-02 15:04:05 MST"),
	}
}

// SetUpdatedFields sets the fields that differ between the two wallets
func (w *Wallet) SetUpdatedFields(srcWallet *Wallet) bool {
	return setUpdatedFields(w, srcWallet)
}

// setUpdatedFields sets the fields in dstStruct that differ from srcStruct
// uses reflect to compare each field and set the updated field
// if the field is a struct, it will recursively call setUpdatedFields on the struct
func setUpdatedFields(dstStruct, srcStruct interface{}) (updated bool) {
	// get the reflect value of the dst struct
	dstStructValue := reflect.ValueOf(dstStruct).Elem()

	// get the reflect value of the src struct
	srcStructValue := reflect.ValueOf(srcStruct).Elem()

	// loop through each field
	for i := 0; i < srcStructValue.NumField(); i++ {
		// get the field
		srcField := srcStructValue.Field(i)
		// if fieldName is BaseNode, skip
		if srcStructValue.Type().Field(i).Name == "BaseNode" {
			continue
		}

		// if the field is a pointer, recursively call setUpdatedFields
		if srcField.Kind() == reflect.Ptr {
			// if the pointer is nil, skip
			if srcField.IsNil() {
				continue
			}

			// recursively call setUpdatedFields
			updated = setUpdatedFields(dstStructValue.Field(i).Interface(), srcField.Interface()) || updated
			continue
		}

		// if the field is array of pointers, recursively call setUpdatedFields
		if srcField.Kind() == reflect.Slice {
			// if the slice is nil, skip
			if srcField.IsNil() {
				continue
			}

			// loop through each element in the slice
			for j := 0; j < srcField.Len(); j++ {
				// recursively call setUpdatedFields
				updated = setUpdatedFields(dstStructValue.Field(i).Index(j).Interface(), srcField.Index(j).Interface()) || updated
			}
			continue
		}

		dstField := dstStructValue.Field(i)

		// if the field is not the same, set the field
		if !reflect.DeepEqual(srcField.Interface(), dstField.Interface()) {
			// set the field
			dstField.Set(srcField)

			// set updated to true
			updated = true
		}
	}

	return updated
}
