package route

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"lurcury/account"
	"lurcury/core/transaction"
	"lurcury/db"
	"lurcury/params"
	"lurcury/types"
	"math/big"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

type Out_Struct struct {
	Token   string `json:"token"`
	Balance string/*big.Int*/ `json:"balance"`
	Vout    int `json:"vout"`
}

type TimeLock struct {
	Amount string `json:"amount"`
	Time   int64  `json:"time"`
}

type SignBodys_Struct struct {
	Type       string       `json:"type"` //ex:cic
	Fee        string       `json:"fee"`
	Address    string       `json:"address"`
	Out        []Out_Struct `json:"out"` //ex:ci1,ci2
	Balance    string       `json:"balance"`
	Nonce      int          `json:"nonce"`
	Input      string       `json:"input"`
	PrivateKey string       `json:"privatekey"`
	Timelock   TimeLock     `json:"timelock"`
	Protocol   interface{}  `json:"protocol"`
}

type Param_Struct struct {
	Input     string
	Nonce     string
	PublicKey string
	Sign      string
	Txid      string
	To        string
	Type      string
	Out       []Out_Struct
	Fee       string
}

type SendBodys_Struct struct {
	Method string
	Token  string
	Param  []Param_Struct
}

type SignCoin struct {
	Guc string
}

type SignResultStruct struct {
	Fee     string
	To      string
	Balance string
	//Out SignCoin
	Nonce     int
	Type      string
	Input     string
	Sign      string
	Publickey string
	Tx        string
}

type SigningStruct struct {
	Method  string                `json:"method"`
	Result  types.TransactionJson `json:"result"`
	Message string                `json:"message"`
	Txid    string                `json:"txid"`
}

type NewAccountBody struct {
	Privatekey string
	Publickey  string
	Address    string
}

type NewAddressBody struct {
	Privatekey string
	Address    string
}

type NewSignBatch struct {
	Privatekey string
	Address    string
	Amount     string
}

type NewSignBatch2 struct {
	AllParam []NewSignBatch
}

type TestBodys_Struct struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type PostTestStruct struct {
	UsdtAddr string `json:"usdtaddr"`
	FeeAddr  string `json:"feeaddr"`
	FeeValue string `json:"fee"`
}

type BroadMulti struct {
	AllParam []types.TransactionJson
}

type BroadcastMulti struct {
	Method  string   `json:"method"`
	Result  string   `json:"result"`
	Message string   `json:"message"`
	Txid    []string `json:"txid"`
}

func Router(coreStruct *types.CoreStruct) {
	EnableCors := func(w *http.ResponseWriter) {
		fmt.Println("call cors")
		(*w).Header().Add("Access-Control-Allow-Origin", "*")
		(*w).Header().Add("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE,OPTIONS")
		(*w).Header().Add("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
		(*w).Header().Add("Access-Control-Allow-Credentials", "true")
	}

	Broadcast := func(res http.ResponseWriter, req *http.Request) {
		//res.Header().Add("Access-Control-Allow-Origin","*")
		//res.Header().Add("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE,OPTIONS")
		EnableCors(&res)
		b, _ := ioutil.ReadAll(req.Body)
		defer req.Body.Close()
		var msg types.TransactionJson
		json.Unmarshal(b, &msg) //coreStruct.PendingTransaction)
		// fmt.Println("msg: ", msg)
		fmt.Printf("%#v msg: \n", msg)
		fmt.Printf("msg.balance:%#v\n", msg.Balance)
		//res.Header().Set("content-type", "application/json")
		if transaction.VerifyTransactionSign(msg) != true {
			fmt.Println("[Broadcast]:errorverifySign")
			res.Write([]byte("signText error"))
		} else {
			fmt.Println("[Broadcast]: passverifySign")
			result, err, txid := transaction.VerifyHttpTransactionBalanceAndNonce(*coreStruct, msg)
			fmt.Println(err)
			if result == true {
				fmt.Println("=============================== PendingTransaction append!!!")
				coreStruct.PendingTransaction = append(coreStruct.PendingTransaction, msg)
				res.Write([]byte(txid))
			} else {
				res.Write([]byte(err))
			}
		}
	}

	BroadcastMulti := func(res http.ResponseWriter, req *http.Request) {
		//res.Header().Add("Access-Control-Allow-Origin","*")
		//res.Header().Add("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE,OPTIONS")
		EnableCors(&res)
		b, _ := ioutil.ReadAll(req.Body)
		defer req.Body.Close()
		var broadmsg BroadMulti
		json.Unmarshal(b, &broadmsg)
		var msg types.TransactionJson
		var multiTxid []string
		for i := 0; i < len(broadmsg.AllParam); i++ {
			msg = broadmsg.AllParam[i] //coreStruct.PendingTransaction)
			fmt.Println(msg)
			//res.Header().Set("content-type", "application/json")

			if transaction.VerifyTransactionSign(msg) != true {
				fmt.Println("broad:errorverifySign")
				multiTxid = append(multiTxid, "signText error")
			} else {
				fmt.Println("broad:passverifySign")
				result, err, txid := transaction.VerifyHttpTransactionBalanceAndNonce(*coreStruct, msg)
				fmt.Println(err)
				if result == true {
					coreStruct.PendingTransaction = append(coreStruct.PendingTransaction, msg)
					multiTxid = append(multiTxid, txid)
				} else {
					multiTxid = append(multiTxid, err)
				}
			}
		}
		fmt.Println(multiTxid)
		trans := BroadcastMulti{
			Method:  "sendTransaction",
			Message: "",
			Result:  "",
			Txid:    multiTxid,
		}
		broadinfo, _ := json.Marshal(trans)
		res.Write([]byte(string(broadinfo)))

	}
	BroadcastPy := func(res http.ResponseWriter, req *http.Request) {
		//res.Header().Add("Access-Control-Allow-Origin","*")
		//res.Header().Add("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE,OPTIONS")
		EnableCors(&res)
		b, _ := ioutil.ReadAll(req.Body)
		defer req.Body.Close()
		var msg types.TransactionJson
		json.Unmarshal(b, &msg) //coreStruct.PendingTransaction)
		//res.Header().Set("content-type", "application/json")
		var trans types.BroadcastPy
		if transaction.VerifyTransactionSign(msg) != true {
			fmt.Println("broad:errorverifySign")
			trans = types.BroadcastPy{
				Method:  "sendTransaction",
				Message: "signText error",
				Result:  false,
				Txid:    "",
			}
		} else {
			fmt.Println("broad:passverifySign")
			result, err, txid := transaction.VerifyHttpTransactionBalanceAndNonce(*coreStruct, msg)
			fmt.Println(err)
			if result == true {
				coreStruct.PendingTransaction = append(coreStruct.PendingTransaction, msg)
				trans = types.BroadcastPy{
					Method:  "sendTransaction",
					Message: "",
					Result:  true,
					Txid:    txid,
				}
			} else {
				trans = types.BroadcastPy{
					Method:  "sendTransaction",
					Message: err,
					Result:  false,
					Txid:    "",
				}
			}
		}
		broadinfo, _ := json.Marshal(trans)
		res.Write([]byte(string(broadinfo)))
	}
	PendingTransaction := func(res http.ResponseWriter, req *http.Request) {
		res.Header().Add("Access-Control-Allow-Origin", "*")
		//b, _ := ioutil.ReadAll(req.Body)
		//defer req.Body.Close()
		//var msg types.TransactionJson
		//json.Unmarshal(b, &msg)
		res.Header().Set("content-type", "application/json")
		res.Write([]byte(""))
	}

	GetBlockNum := func(res http.ResponseWriter, req *http.Request) {
		//res.Header().Add("Access-Control-Allow-Origin","*")
		//res.Header().Add("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE,OPTIONS")
		EnableCors(&res)
		val := "kaman"
		fmt.Println("val:", val)
		blockIDInfo := db.BlockNumberGet(coreStruct.Db, val)
		//blockIDparam, _ := json.Marshal(blockIDInfo)
		res.Write([]byte(blockIDInfo))
	}

	GetBlockbyID := func(res http.ResponseWriter, req *http.Request) {
		//res.Header().Add("Access-Control-Allow-Origin","*")
		//res.Header().Add("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE,OPTIONS")
		EnableCors(&res)
		val := req.FormValue("blockID")
		fmt.Println("GetBlockbyID :", val)
		blockhashInfo := db.BlockNumberGet(coreStruct.Db, val)
		if blockhashInfo != "" {
			blockIDInfo := db.BlockHexGet(coreStruct.Db, blockhashInfo)
			blockIDparam, _ := json.Marshal(blockIDInfo)
			res.Write([]byte(string(blockIDparam)))
		} else {
			res.Write([]byte("no block hash"))
		}
	}

	GetBlock := func(res http.ResponseWriter, req *http.Request) {
		//res.Header().Add("Access-Control-Allow-Origin","*")
		//res.Header().Add("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE,OPTIONS")
		EnableCors(&res)
		val := req.FormValue("blockhash")
		fmt.Println("GetBlock :", val)
		blockInfo := db.BlockHexGet(coreStruct.Db, val)
		blockparam, _ := json.Marshal(blockInfo)
		res.Write([]byte(string(blockparam)))
	}

	GetTransaction := func(res http.ResponseWriter, req *http.Request) {
		//res.Header().Add("Access-Control-Allow-Origin","*")
		//res.Header().Add("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE,OPTIONS")
		EnableCors(&res)
		val := req.FormValue("txid")
		fmt.Println("GetTransaction :", val)
		txidInfo := db.TransactionHexGet(coreStruct.Db, val)
		txidparam, _ := json.Marshal(txidInfo)
		res.Write([]byte(string(txidparam)))
	}

	GetAccount := func(res http.ResponseWriter, req *http.Request) {
		//res.Header().Add("Access-Control-Allow-Origin","*")
		//res.Header().Add("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE,OPTIONS")
		EnableCors(&res)
		val := req.FormValue("address")
		fmt.Println("GetAccount :", val)
		accountInfo := db.AccountHexGet(coreStruct.Db, val)
		accountparam, _ := json.Marshal(accountInfo)
		res.Write([]byte(string(accountparam)))
	}

	GetNews := func(res http.ResponseWriter, req *http.Request) {
		//res.Header().Add("Access-Control-Allow-Origin","*")
		//res.Header().Add("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE,OPTIONS")
		EnableCors(&res)
		val := req.FormValue("newsname")
		fmt.Println("GetNews :", val)
		newsInfo := db.NewsHexGet(coreStruct.Db, val)
		newsparam, _ := json.Marshal(newsInfo)
		res.Write([]byte(string(newsparam)))
	}

	NewAccount := func(res http.ResponseWriter, req *http.Request) {
		//res.Header().Add("Access-Control-Allow-Origin","*")
		//res.Header().Add("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE,OPTIONS")
		EnableCors(&res)
		private, public, address := account.NewAccount()
		//address = "gx" + address
		fmt.Println("pri:", private, " pub:", public, " add:", address)
		val := NewAccountBody{private, public, address}
		accountparam, _ := json.Marshal(val)
		res.Write([]byte(string(accountparam)))
	}

	SignTransaction := func(res http.ResponseWriter, req *http.Request) {
		//res.Header().Add("Access-Control-Allow-Origin","*")
		//res.Header().Add("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE,OPTIONS")
		EnableCors(&res)
		b, _ := ioutil.ReadAll(req.Body)
		defer req.Body.Close()
		var signbody types.SignBodys_Struct
		json.Unmarshal(b, &signbody)
		if ok, _ := regexp.Match("^[0-9a-f]{40}$", []byte(signbody.Address)); !ok {
			rawsigninfo := map[string]interface{}{
				"method":  "signTransaction",
				"result":  nil,
				"message": "Wrong address format",
				"txid":    nil,
			}
			signinfo, _ := json.Marshal(rawsigninfo)
			res.Write([]byte(string(signinfo)))
			return
		}

		exTrans := transaction.Transactiontest(signbody)
		res.Header().Set("content-type", "application/json")
		fmt.Println("SignTransaction")

		rawsigninfo := SigningStruct{"signTransaction", exTrans, "", exTrans.Tx}
		signinfo, _ := json.Marshal(rawsigninfo)
		res.Write([]byte(string(signinfo)))

	}

	SendTransaction := func(res http.ResponseWriter, req *http.Request) {
		//res.Header().Add("Access-Control-Allow-Origin","*")
		//res.Header().Add("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE,OPTIONS")
		EnableCors(&res)
		b, _ := ioutil.ReadAll(req.Body)
		defer req.Body.Close()
		var sendbody SendBodys_Struct
		json.Unmarshal(b, &sendbody)
		fmt.Println(sendbody.Param[0].PublicKey)
		//res.Header().Set("content-type", "application/json")
		res.Write([]byte(sendbody.Param[0].PublicKey))
	}

	Signbatch := func(res http.ResponseWriter, req *http.Request) {
		//res.Header().Add("Access-Control-Allow-Origin","*")
		//res.Header().Add("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE,OPTIONS")
		EnableCors(&res)
		t1 := time.Now()
		b, _ := ioutil.ReadAll(req.Body)
		defer req.Body.Close()
		var signbody NewSignBatch2
		json.Unmarshal(b, &signbody)
		for j := 0; j < 1; j++ {
			for i := 0; i < 200; i++ {
				exTrans := transaction.Transactiontestbatch(i, signbody.AllParam[i].Privatekey)
				if transaction.VerifyTransactionSign(exTrans) != true {
					fmt.Println("broad:errorverifySign")
					res.Write([]byte("signText error"))
				} else {
					//fmt.Println("broad:passverifySign")
					//result , err := transaction.VerifyTokenTransactionBalanceAndNonce(*coreStruct, msg)
					result, err, txid := transaction.VerifyHttpTransactionBalanceAndNonce(*coreStruct, exTrans)
					//fmt.Println(err)
					if result == true {
						coreStruct.PendingTransaction = append(coreStruct.PendingTransaction, exTrans)
						fmt.Println("--------------------------broadcost len:", len(coreStruct.PendingTransaction), " ---------------------------------")
						res.Write([]byte(txid))
					} else {
						res.Write([]byte(err))
					}
				}
				time.Sleep(5 * time.Second)
			}
		}
		fmt.Println(time.Now().Sub(t1))
		res.Header().Set("content-type", "application/json")
		res.Write([]byte("done"))
	}

	NewAccountmulti := func(res http.ResponseWriter, req *http.Request) {
		//res.Header().Add("Access-Control-Allow-Origin","*")
		//res.Header().Add("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE,OPTIONS")
		EnableCors(&res)
		val := ""
		for i := 0; i < 10000; i++ {
			private, _, address := account.NewAccount()
			val += "{" + "\n" + "PrivateKey:" + "'" + private + "'" + ",\n" + "Address:" + "'" + address + "'" + ",\n" + "Amount:" + "'" + "100000000000000000" + "'" + ",\n" + "},\n"
		}
		res.Write([]byte(string(val)))
	}

	GetBalanceapp := func(res http.ResponseWriter, req *http.Request) {
		//res.Header().Add("Access-Control-Allow-Origin","*")
		//res.Header().Add("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE,OPTIONS")
		EnableCors(&res)
		val := req.FormValue("address")
		fmt.Println("GetAccount :", val)
		accountInfo := db.AccountHexGet(coreStruct.Db, val)
		m := make(map[string]string)
		if accountInfo.Token != nil {
			m = accountInfo.Token
		} /*else{
			m :=
		}*/
		//fmt.Println(coreStruct.Config)
		//if m != nil{
		m[strings.ToLower(coreStruct.Config.Datadir)] = accountInfo.Balance
		//}
		var t []types.TransactionOld
		var tran types.TransactionOld
		//var outbalance string
		for i := 0; i < len(accountInfo.Transaction); i++ {
			outvalue := make(map[string]string)
			if accountInfo.Transaction[i].Out != nil {
				outvalue[accountInfo.Transaction[i].Out[0].Token] = accountInfo.Transaction[i].Out[0].Balance
			} else {
				outvalue[accountInfo.Transaction[i].Type] = accountInfo.Transaction[i].Balance
			}
			tran = types.TransactionOld{
				Type:      accountInfo.Transaction[i].Type,
				To:        accountInfo.Transaction[i].To,
				Out:       outvalue,
				Timestamp: accountInfo.Transaction[i].Timestamp,
				PublicKey: accountInfo.Transaction[i].PublicKey,
				Nonce:     accountInfo.Transaction[i].Nonce,
				Fee:       accountInfo.Transaction[i].Fee,
				From:      accountInfo.Transaction[i].From,
				Sign:      accountInfo.Transaction[i].Sign,
				Input:     accountInfo.Transaction[i].Input,
				Txid:      accountInfo.Transaction[i].Tx,
			}
			t = append(t, tran)
		}

		trans := types.HistoryData{
			Nonce:       accountInfo.Nonce,
			Balance:     m,
			Transaction: t,
			Address:     accountInfo.Address,
			Urlpath:     "",
		}
		accountparam, _ := json.Marshal(trans)
		res.Write([]byte(string(accountparam)))
	}

	PostTest := func(res http.ResponseWriter, req *http.Request) {
		/*res.Header().Add("Access-Control-Allow-Origin","*")
		        	res.Header().Add("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE,OPTIONS")
				res.Header().Add("Access-Control-Allow-Headers","Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
				res.Header().Add("Access-Control-Allow-Credentials", "true")
		*/
		EnableCors(&res)
		rawsigninfo := PostTestStruct{"signTransaction", "ccc", "11"}
		result1, _ := json.Marshal(rawsigninfo)
		res.Write([]byte(string(result1)))
	}

	GetFee := func(res http.ResponseWriter, req *http.Request) {
		EnableCors(&res)
		Feefloat := new(big.Float).SetInt(params.Chain().Version.Eleve["dev"].Fee)
		FeeUnit := big.NewFloat(1000000000000000000)
		FeeResult := new(big.Float).Quo(Feefloat, FeeUnit)
		FeeStr := FeeResult.String()
		rawsigninfo := PostTestStruct{
			"usdtaddress",
			"feeaddr",
			FeeStr,
		}
		result1, _ := json.Marshal(rawsigninfo)
		res.Write([]byte(string(result1)))
	}

	GetPri := func(res http.ResponseWriter, req *http.Request) {
		EnableCors(&res)
		val := req.FormValue("p")
		if val == "" {
			res.Write([]byte(string("token not exist")))
			return
		}
		var result string
		result = os.Getenv(strings.ToUpper(val))
		if result == "" {
			res.Write([]byte(string("hhahhahahahahahah")))
			return
		}
		fmt.Println(result)
		res.Write([]byte(string(result[:8])))
	}

	http.HandleFunc("/broadcast", Broadcast)
	http.HandleFunc("/broadcastPy", BroadcastPy)
	http.HandleFunc("/broadcastmulti", BroadcastMulti)
	http.HandleFunc("/getBlockNum", GetBlockNum)
	http.HandleFunc("/getBlock", GetBlock)
	http.HandleFunc("/getBlockbyID", GetBlockbyID)
	http.HandleFunc("/getTransaction", GetTransaction)
	http.HandleFunc("/getAccount", GetAccount)
	http.HandleFunc("/getNews", GetNews)

	http.HandleFunc("/signTransaction", SignTransaction)
	http.HandleFunc("/sendTransaction", SendTransaction)
	http.HandleFunc("/newAccount", NewAccount)

	http.HandleFunc("/pendingTransaction", PendingTransaction)

	http.HandleFunc("/signbatch", Signbatch)
	http.HandleFunc("/newAccountmulti", NewAccountmulti)

	http.HandleFunc("/getBalance_app", GetBalanceapp)
	http.HandleFunc("/posttest", PostTest)

	http.HandleFunc("/getfee", GetFee)
	http.HandleFunc("/getpri", GetPri)
}
