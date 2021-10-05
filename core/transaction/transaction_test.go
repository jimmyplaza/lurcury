package transaction

import (
	"encoding/hex"
	"fmt"
	"lurcury/crypto"
	"lurcury/crypto/eddsa"
	secp "lurcury/crypto/secp256k1"
	"lurcury/types"

	//"lurcury/account"
	//"lurcury/crypto"
	//"lurcury/db"
	//"lurcury/types"
	//	"lurcury/core/transaction"
	//"math/big"
	//"reflect"
	//"container/list"
	"testing"
	//"encoding/json"
	//"time"
)

/*
func TestLock(t *testing.T){
        core_arg := &types.CoreStruct{}
        core_arg.Db = db.OpenDB("../../dbdata")
	ex := ExpLockTransaction()
	fmt.Println(ex)
	x,_:=LockProtocol(*core_arg, ex)
	fmt.Println(x)
}
*/
/*
func TestPro(t *testing.T){
	ex := ExpBnnTransaction()
	fmt.Println(ex.Protocol.(types.NewsStation))
}
*/

/*
func TestTop(t *testing.T){
	x := NewTran("cic")
	//y,_:=hex.DecodeString("219a634773d787cfbaf1e5c915d56b14be2a3695ed8e46bbeb01573bf211d0ef8773580834eb42a2f2ee856b029a88dfee639e27f08b1e0235f8eb04eecf4089")
        y,_:=hex.DecodeString("A665A45920422F9D417E4867EFDC4FB8A04A1F3FFF1FA07E998E86F7F7A27AE3")
	fmt.Println(EncodeForSign(x))
	fmt.Println(SignTransaction(y,x).Sign)
	fmt.Println(VerifyTransactionSign(SignTransaction(y,x)))

}
*/

func TestPublicKey(t *testing.T) {
	var privateKey = "1c40de38b5d888b51aaf5871320f3ac77df4ba2968a1b8c38ff6b6e0986426bb"
	Pri, _ := hex.DecodeString(privateKey)
	var strpublicKey string
	var publicKey []byte
	//switch SignBody.Crypto {
	//case "cic":
	publicKey = secp.SecpKeyToPublicKey(Pri)
	//default:
	//	publicKey = secp.SecpKeyToPublicKey(Pri)
	//}
	strpublicKey = hex.EncodeToString(publicKey)
	fmt.Println("strpublicKey:" + strpublicKey)

	out := []types.TransactionOut{
		{
			Token:   "usdtn",
			Balance: "1000000000",
			Vout:    0,
		},
	}
	trans := types.TransactionJson{
		Balance:   "0",
		Out:       out,
		To:        "11a0a990785b12d2bc09d35ac5a31a516ae9b77e",
		Nonce:     43,
		Fee:       "0",
		Type:      "bolt",
		Input:     "usdtn$mcc$0.3333333333333333",
		Crypto:    "cic",
		PublicKey: "042eed63859378e244d25a6882e0c4fa46575fd22d033d3e90003fe2ffefd7668e3e1ca09f7905d69cdc511e6e511b26e15236cc2433e8918bcd5533246d0e3e07",
		Tx:        "134ad62ba88b499b024d7e781a8008e045b4af83ec62ae3b680c7bf06eb2939e",
		Sign:      "3538343839393238363431313632383532363135353237303634303236343434343131323533343537373232303333363032353238333334343531303832353532313032383434303733393937783439303631363036393534353539383637353739333834323032343836303635383133373637323139313233363237323033333236313933383130333133393032333334303437373731373037",
	}

	encodeSign := EncodeForSign(trans)
	fmt.Println("encodeSign:" + encodeSign)
	tx := hex.EncodeToString(crypto.Sha256([]byte(EncodeForSign(trans))))
	fmt.Println("tx:" + tx)
	re := verifyTransactionSign(trans)
	fmt.Println(re)
}

func verifyTransactionSign(Transaction types.TransactionJson) bool {
	pub, _ := hex.DecodeString(Transaction.PublicKey)
	fmt.Println("[VerifyTransactionSign] Transaction.PublicKey: ", Transaction.PublicKey)
	fmt.Printf("%#v ", Transaction)
	msg := crypto.Keccak256([]byte(EncodeForSign(Transaction)))
	if Transaction.Crypto == "a64" {
		fmt.Println("Transaction.Crypto is a64!!")
		msg = crypto.Sha256([]byte(EncodeForSign_Simple(Transaction)))
	}
	sign, _ := hex.DecodeString(Transaction.Sign)
	fmt.Println("[VerifyTransactionSign] hex.EncodeToString(msg): ", hex.EncodeToString(msg))
	//fmt.Println("msg",EncodeForSign_Simple(Transaction))
	var re bool
	switch Transaction.Crypto {
	case "cic":
		re = secp.SecpVerify(pub,
			msg,
			sign,
		)
	case "secp256k1":
		re = secp.SecpVerify(pub,
			msg,
			sign,
		)
	case "eddsa":
		re = eddsa.EddsaVerify(pub,
			msg,
			sign,
		)
	case "a64":
		re = secp.SecpVerify2(pub,
			msg,
			sign,
		)
	default:
		re = secp.SecpVerify(pub,
			msg,
			sign,
		)
	}
	return re
}

//func TestSecp2(t *testing.T){
//        trans := NewTran("a64")
//	trans.PublicKey = "04be686ed7f0539affbaf634f3bcc2b235e8e220e7be57e9397ab1c14c39137eb43705125aac75a865268ef33c53897c141bd092cf4d1a306b2a57e37e1386826d"
//        //testString := EncodeForSign_Simple(trans)
//	key,_:=hex.DecodeString("A665A45920422F9D417E4867EFDC4FB8A04A1F3FFF1FA07E998E86F7F7A27AE3")
//	sign := SignTransaction(key, trans)
//	fmt.Println("sign result:", sign.Sign)
//        fmt.Println(VerifyTransactionSign(SignTransaction(key, trans)))
//
//}

/*
func TestTop2(t *testing.T){
        trans := types.TransactionJson{
                Balance: "1",
                To: "a3d5b73a8e19e763df8ed9eb3e97c78958d440fb",
                Nonce: 1,
                Fee: "1",
                Type: "aaa",
                Input: "daa11234",
                Crypto: "xpe",
		PublicKey: "8773580834eb42a2f2ee856b029a88dfee639e27f08b1e0235f8eb04eecf4089",
        }
        y,_:=hex.DecodeString("219a634773d787cfbaf1e5c915d56b14be2a3695ed8e46bbeb01573bf211d0ef8773580834eb42a2f2ee856b029a88dfee639e27f08b1e0235f8eb04eecf4089")
        fmt.Println("加縮後格式：",EncodeForSign_Simple(trans))
        fmt.Println("簽章後結果",SignTransaction(y,trans).Sign)
        fmt.Println("pub:",SignTransaction(y,trans).PublicKey)
        fmt.Println("驗證後結果:",VerifyTransactionSign(SignTransaction(y,trans)))
	re := SignTransaction(y,trans)
	fmt.Println(re.Sign)
	re.Sign = "7408cef4eac60352acccf8019b5de49df4da808ff1f845e9082bad27a377b43afc0e2a20ead894ab4cdd7a488bc8a0c6cb073025bd96f956ac38d57344b23b01"
	fmt.Println("驗證後結果:",VerifyTransactionSign(re))
}
*/
/*
func TestString(t *testing.T){
	ex := ExpBnnTransaction()
	//fmt.Println(ex)
	a,_ := json.Marshal(ex)
	var b types.TransactionJson
	json.Unmarshal(a,&b)
	//fmt.Println(b.Protocol)
	d := b.Protocol.(map[string]interface{})
	//fmt.Println(d["picture"])
	e := d["picture"].([]interface{})
	x := e[0].(map[string]interface{})
	for key, value := range e {
		fmt.Println("Key:", key, "Value:", value)
	}
	fmt.Println(x["title"])
	//v.(struct{Id int}).Id
	//var d types.NewsStation
	//json.Unmarshal(b.Protocol.(types.NewsStation),&d)
	//var c map[string]interface{}
	//json.Unmarshal([]byte(b.Protocol), &c)
	//fmt.Println(b.Protocol["name"])
}
*/
/*
func TestTime(t *testing.T){
	timestamp := time.Now().Unix()
	fmt.Println("timestamp:",timestamp)

	var timeint64 int64
	timeint64=1534329801
	tm := time.Unix(timeint64,0)
	tm2 := time.Unix(1534329805,0)
	fmt.Println("tm:",tm)
	fmt.Println("tm2:",tm2)
}
*/
/*
func TestNil(t *testing.T){
    var testbody types.TransactionJson
    if (testbody.Protocol == nil){
    	fmt.Println("protocol nil:",testbody.Protocol)
    }else{
    	fmt.Println("protocol nonil:",testbody.Protocol)
    }
}
*/
/*
func TestDecode(t *testing.T){
	//NodeinfoDecode("444wx00300600100145678911111111",3)
	ex := ExpBnnTransaction()
	re := BnnTransactionToStation(ex)
	fmt.Println(re)
}
*/
/*
func TestBigTransaction(t *testing.T){
        core_arg := &types.CoreStruct{}
        core_arg.Db = db.OpenDB("../../dbdata")

	d := ExpTokenTransaction()
	fmt.Println(d.Input[0:3])
	TransactionProtocol(*core_arg, d)
}
*/
/*
func TestBigTransaction(t *testing.T){
        fromBalance := new(big.Int)
        toBalance := new(big.Int)
        transBalance := new(big.Int)
        fromBalance.SetString("10000000000000000000000000000",10)
        toBalance.SetString("8000000000000000000000000000",10)
        transBalance.SetString("6000000000000000000000000000",10)
        if(fromBalance.Cmp(transBalance)>=0){
                fromBalance.Sub(fromBalance, transBalance)
                toBalance.Add(toBalance, transBalance)
        }
	fmt.Println(fromBalance)
	fmt.Println(toBalance.String())
}
/*
func TestPendingTransaction(t *testing.T){

        core_arg := &types.CoreStruct{}
        core_arg.Db = db.OpenDB("../../dbdata")
        bb := ExpTransaction()
        core_arg.PendingTransaction = append(core_arg.PendingTransaction, bb)
        core_arg.PendingTransaction = append(core_arg.PendingTransaction, bb)
	fmt.Println("initTransaction:",core_arg.PendingTransaction)
	deletp := DeletPendingTransaction(*core_arg,0)
	fmt.Println("deleteTransaction:",deletp)
	fmt.Println(core_arg.PendingTransaction)
}
*/

/*
func TestTransaction(t *testing.T){
        check := func(f string, got, want interface{}) {
                if !reflect.DeepEqual(got, want) {
                        t.Errorf("%s mismatch: got %v, want %v", f, got, want)
                }
        }

	core_arg := &types.CoreStruct{}
        core_arg.Db = db.OpenDB("../dbdata")

	//初始化金額

        account_tmp := account.Account_exp()
	//fmt.Println(account_tmp)
	db.AccountHexPut(core_arg.Db, account_tmp.Address, account_tmp)

	pp := Transactiontestbatch(0)//ExpTransaction()
	fmt.Println("ppgo:",pp)
	fmt.Println("sign verify test:",VerifyTransactionSign(pp))
	//fmt.Println("test token amount:", pp.Out[0].Token)
	fmt.Println("from address test:",crypto.KeyToAddress_hex(pp.PublicKey))
	//fmt.Println("Nonce:",account_tmp.Nonce)

	fmt.Println("sign verify test:",VerifyTransactionSign(pp))
	fmt.Println("pp:",pp)

	fmt.Println(pp)
	m1, m2 := VerifyTokenTransactionBalanceAndNonce(*core_arg, pp)
	fmt.Println("verify balance and nonce:",m1)
	fmt.Println("balanceresult:",m2)



	a3 := db.AccountHexGet(core_arg.Db, "264411884d6d2aca8ca2d2a77c9dc95ffdcee521")
	fmt.Println("test for verify balance and nonce result:",a3)
	fmt.Println(pp.PublicKey)
	a4 := db.AccountHexGet(core_arg.Db, crypto.KeyToAddress_hex(pp.PublicKey))
	fmt.Println(a4)

	check("go","123","123")

}

*/
