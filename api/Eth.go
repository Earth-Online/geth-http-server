package api

import (
	"math/big"
	"fmt"
)

func init() {

}

func EthGetBalance(address string,tag string) (*big.Int,error) {

	bi := big.NewInt(0)
   arg := map[string]string {
   	  "address":address,
   	  "tag":tag,
   }
	data,err := Call(AccountModule,"balance",arg)
	if err != nil {
		return nil,err
	}

	num,_ := bi.SetString(data.Result.(string),10)
	return num,nil

	}

func EthGetTransactionByHash(txHash string) (*TransactionObject, error) {

	arg := map[string]string {
		"txhash":txHash,
	}
	resp, err :=  Call(ProxyModule,"eth_getTransactionByHash",arg)
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	answer := new(TransactionObject)
	err = MapToObject(resp.Result, answer)
	if err != nil {
		return nil, err
	}
	return answer, nil
}
