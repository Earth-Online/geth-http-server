package api

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
)

func Update(address string,password string, newstring string) error {
	AccountAddress := common.HexToAddress(address)
	account := accounts.Account{
		Address: AccountAddress,
		URL:     accounts.URL{Scheme:keystore.KeyStoreScheme,Path:""},
	}
	//accounts.URL{Scheme: keystore.KeyStoreScheme, Path: (keyFileName(key.From))}}
	return key.Update(account,password,newstring)
}
