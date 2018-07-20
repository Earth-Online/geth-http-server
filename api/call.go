package api

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"math/rand"
	"time"
	"strings"
	"web/conf"
)

var   (
	ApiKey []string
	key *keystore.KeyStore
)

const (
	ApiUrl = "https://api.etherscan.io"
	ProxyModule = "proxy"
	AccountModule = "account"
	StatsModule = "stats"


)

func init() {
	ApiKey = strings.Split(conf.Api.Key("ethscan_apikey").String(),",")
    key = keystore.NewKeyStore(conf.Rpc.Key("KeystoreDir").String(),keystore.StandardScryptN,keystore.StandardScryptP)
	rand.Seed(time.Now().Unix())
}

func Call(module string, action string,arg map[string]string) (*JSON2Response,error) {
	var urlArg string
    for key,value := range arg {
    	urlArg = urlArg + fmt.Sprintf("&%s=%s",key,value)
	}
	apikey := ApiKey[rand.Intn(len(ApiKey)-1)]
	url := fmt.Sprintf("%s/api?module=%s&action=%s%s&apikey=&%s",ApiUrl,module,action,urlArg,apikey)

	response,err  := http.Get(url)

	if err != nil {
		return nil,err
	}


	if response.StatusCode  != 200 {
		return  nil, err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return  nil, err
	}

	jResp := new(JSON2Response)

	err = json.Unmarshal(body, jResp)
	if err != nil {
		return nil, err
	}

	return jResp,nil

}




