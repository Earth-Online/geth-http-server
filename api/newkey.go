package api

import (
	"github.com/ethereum/go-ethereum/crypto"
	"encoding/hex"
)

func NewKey(password string) (string,string,error) {
	PrivateKey,_,err := key.NewAccountAndKey(password)
	if err != nil {
		return  "","",err
	}
	return  hex.EncodeToString(crypto.FromECDSA(PrivateKey.PrivateKey)),PrivateKey.Address.String(),nil
}
