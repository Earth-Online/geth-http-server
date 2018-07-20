package api

import (
	"io/ioutil"
	"github.com/ethereum/go-ethereum/accounts/keystore"
)

func KeyEncryptDecrypt(file string,password string)(*keystore.Key,error){
	keyjson, err := ioutil.ReadFile(file)
	if err != nil {
		return nil,err
	}
	key, err := keystore.DecryptKey(keyjson, password)
	if err != nil {
		return nil,err
	}
	return key,nil
}

