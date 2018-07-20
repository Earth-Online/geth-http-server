package api

import "testing"

func TestEthGetBalance(t *testing.T) {
	ret,err := EthGetBalance("0xde0b295669a9fd93d5f28d9ec85e40f4cb697bae","latest")
	if err != nil {
		t.Error(err)
	}
	t.Log(ret)
	}


func TestEthGetTransactionByHash(t *testing.T) {
	ret,err := EthGetTransactionByHash("0x1e2910a262b1008d0616a0beb24c1a491d78771baa54a33e66065e03b1f46bc1")
	if err != nil {
		t.Error(err)
	}
	t.Log(ret.To)
	t.Log(ret.Input)
	t.Log(ret.Value)
}

func TestEthprice(t *testing.T) {
	ret,err := Ethprice()
	if err != nil {
		t.Error(err)
	}
	t.Log(ret.Ethusd)
}


func TestUpdate(t *testing.T) {

  err := Update("0x2e4d4c18e255fd4c14cf3e5f0e0cae82236cf9a3","123456","****")
	if err != nil {
		t.Error(err)
	}
  }
func TestNewKey(t *testing.T) {
	PrivateKey,address,err := NewKey("123456")
	if err != nil {
		t.Error(err)
	}
	t.Logf(PrivateKey)
	t.Logf(address)
}
func TestQueryTransactions(t *testing.T) {
	ret,err :=QueryTransactions("0x018404ab6bb0790c646dda5e939f6aacf8805dc3")
	if err != nil {
		t.Error(err)
	}
	t.Logf(reflect.ValueOf(ret).String())
	t.Logf(reflect.TypeOf(ret).String())
	t.Logf(fmt.Sprint(ret))
}
