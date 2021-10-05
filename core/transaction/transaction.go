package transaction

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"lurcury/crypto"
	eddsa "lurcury/crypto/eddsa"
	secp "lurcury/crypto/secp256k1"
	"lurcury/db"
	"lurcury/types"
	"math/big"
	"strconv"
	"strings"
	"time"
	//"reflect"
)

func TransactionModify(BlockHash string, Transaction types.TransactionJson) types.TransactionJson {
	Transaction.BlockHash = BlockHash
	var address string
	switch Transaction.Crypto {
	case "cic":
		address = crypto.CICKeyToAddress_hex(Transaction.PublicKey)
	case "secp256k1":
		address = crypto.CICKeyToAddress_hex(Transaction.PublicKey)
	case "eddsa":
		address = crypto.KeyToAddress_hex(Transaction.PublicKey)
	default:
		address = crypto.CICKeyToAddress_hex(Transaction.PublicKey)
	}
	Transaction.From = address
	Transaction.Version = "sDAGraph"
	Transaction.Tx = hex.EncodeToString(crypto.Sha256([]byte(EncodeForSign(Transaction))))

	//if(Transaction.Crypto=="a64"){
	//        Transaction.Tx = hex.EncodeToString(crypto.Sha256([]byte(EncodeForSign_Simple(Transaction))))
	//}
	return Transaction
}

func NewTokenTransaction(Balance string, To string, Token string, TokenBalance string /*big.Int*/, Nonce int, Fee string /*big.Int*/, Type string, Input string) types.TransactionJson {
	out := []types.TransactionOut{
		{Balance: TokenBalance, Token: Token, Vout: 0},
	}

	trans := types.TransactionJson{
		Balance: Balance,
		Out:     out,
		To:      To,
		Nonce:   Nonce,
		Fee:     Fee,
		Type:    Type,
		Input:   Input,
	}

	return trans
}

func NewTransaction(Balance string, To string, Nonce int, Fee string, Input string, Type string) types.TransactionJson {
	trans := types.TransactionJson{
		Balance: Balance,
		To:      To,
		Nonce:   Nonce,
		Fee:     Fee,
		Type:    Type,
		Input:   Input,
	}
	return trans
}

func SignTransaction(pri []byte, Transaction types.TransactionJson) types.TransactionJson {
	Transaction.Tx = hex.EncodeToString(crypto.Keccak256([]byte(EncodeForSign(Transaction))))
	fmt.Println(EncodeForSign_Simple(Transaction))
	fmt.Println("tag:", hex.EncodeToString(crypto.Sha256([]byte(EncodeForSign_Simple(Transaction)))))
	var sign []byte
	var publicKey []byte
	fmt.Println("crypto:", Transaction.Crypto)
	switch Transaction.Crypto {
	case "eddsa":
		sign = eddsa.EddsaSign(pri, crypto.Keccak256([]byte(EncodeForSign(Transaction))))
		publicKey = eddsa.EddsaKeyToPublicKey(pri)
	case "secp256k1":
		sign = secp.SecpSign(pri, crypto.Keccak256([]byte(EncodeForSign(Transaction))))
		publicKey = secp.SecpKeyToPublicKey(pri)
	case "cic":
		sign = secp.SecpSign(pri, crypto.Keccak256([]byte(EncodeForSign(Transaction))))
		publicKey = secp.SecpKeyToPublicKey(pri)
	case "a64":
		sign = secp.SecpSign2(pri, crypto.Sha256([]byte(EncodeForSign_Simple(Transaction))))
		//fmt.Println("kec:",EncodeForSign_Simple(Transaction))
		publicKey = secp.SecpKeyToPublicKey(pri)
	default:
		sign = eddsa.EddsaSign(pri, crypto.Keccak256([]byte(EncodeForSign(Transaction))))
		publicKey = eddsa.EddsaKeyToPublicKey(pri)
	}
	Transaction.Sign = hex.EncodeToString(sign)
	Transaction.PublicKey = hex.EncodeToString(publicKey)
	Transaction.Tx = hex.EncodeToString(crypto.Keccak256([]byte(EncodeForSign(Transaction))))
	fmt.Println(Transaction.Tx)
	return Transaction
}

/*
func SignTransaction_hex(prik string, Transaction types.TransactionJson)(types.TransactionJson){
	pri,_ := hex.DecodeString(prik)
	Transaction.Tx = string(crypto.Keccak256([]byte(EncodeForSign(Transaction))))
        sign := eddsa.EddsaSign(pri, crypto.Keccak256([]byte(EncodeForSign(Transaction))))
        r := eddsa.EddsaKeyToPublicKey(pri)
        Transaction.Sign = hex.EncodeToString(sign)
        Transaction.PublicKey = hex.EncodeToString(r)
        return Transaction
}
*/

//crypto待加
func EncodeForSign(Transaction types.TransactionJson) string {
	if Transaction.Crypto == "a64" {
		return EncodeForSign_Simple(Transaction)
	}
	from := "wx" + StringTransactionEncode(Transaction.PublicKey, 42)
	to := "gx" + StringTransactionEncode(Transaction.To /*[2:]*/, 42)
	nonce_tmp := strconv.FormatInt(int64(Transaction.Nonce), 16)
	nonce := "hx" + StringTransactionEncode(nonce_tmp, 32)
	fee := "ix" + StringTransactionEncode(Transaction.Fee, 32)
	//typ_tmp:= hex.EncodeToString([]byte(Transaction.Type))
	typ := "kx" + StringTransactionEncode(Transaction.Type /*typ_tmp*/, 8)
	//input_tmp_str := hex.EncodeToString([]byte(Transaction.Input))
	input_tmp := strconv.FormatInt(int64(len(Transaction.Input /*input_tmp_str*/)), 16)
	input := "lx" + StringTransactionEncode(input_tmp, 3) + Transaction.Input /*string(input_tmp_str[:])*/
	balance := "sx" + StringTransactionEncode(Transaction.Balance /*[2:]*/, 32)

	proto, _ := json.Marshal(Transaction.Protocol)
	protocol := "tx" + StringTransactionEncode(string(proto), 8)

	outResult := ""
	token := ""
	tokenBalance := ""
	vout := ""
	for i := 0; i < len(Transaction.Out); i++ {
		token_tmp := hex.EncodeToString([]byte(Transaction.Out[i].Token))
		token = StringTransactionEncode(string(token_tmp), 8)
		//balance_tmp   := strconv.FormatInt(int64(Transaction.Out[i].Balance.Uint64()), 16)
		tokenBalance = StringTransactionEncode(Transaction.Out[i].Balance, 32) //balance_tmp, 32)
		vout_tmp := strconv.Itoa(Transaction.Out[i].Vout)
		vout = StringTransactionEncode(vout_tmp, 3)
		outResult = outResult + "px" + vout + token + tokenBalance
	}
	re := from + to + nonce + fee + typ + outResult + input + balance + protocol
	return re
}

func EncodeForSign_Simple(Transaction types.TransactionJson) string {

	from := Transaction.PublicKey
	to := StringTransactionEncode(Transaction.To, 40)
	nonce_tmp := strconv.FormatInt(int64(Transaction.Nonce), 10)
	nonce := StringTransactionEncode(nonce_tmp, 10)
	fee := StringTransactionEncode(Transaction.Fee, 40)
	typ := StringTransactionEncode(Transaction.Type, 10)
	crypto := StringTransactionEncode(Transaction.Crypto, 10)
	input := Transaction.Input
	balance := StringTransactionEncode(Transaction.Balance, 40)
	/*
		fmt.Println("from:", from)
		fmt.Println("to:", to)
		fmt.Println("nonce:", nonce)
		fmt.Println("fee:", fee)
		fmt.Println("typ:", typ)
		fmt.Println("input:", input)
		fmt.Println("balance:", balance)
	*/
	re := to + from + balance + nonce + fee + typ + crypto + input
	return re
}

func EncodeTransaction(Transaction types.TransactionJson) string {
	from := "wx" + StringTransactionEncode(Transaction.PublicKey, 42)
	to := "gx" + StringTransactionEncode(Transaction.To /*[2:]*/, 42)
	nonce_tmp := strconv.FormatInt(int64(Transaction.Nonce), 16)
	nonce := "hx" + StringTransactionEncode(nonce_tmp, 32)
	fee := "ix" + StringTransactionEncode(Transaction.Fee, 32)
	//typ_tmp:= hex.EncodeToString([]byte(Transaction.Type))
	typ := "kx" + StringTransactionEncode(Transaction.Type /*typ_tmp*/, 8)
	//input_tmp_str := hex.EncodeToString([]byte(Transaction.Input))
	input_tmp := strconv.FormatInt(int64(len(Transaction.Input /*input_tmp_str*/)), 16)
	input := "lx" + StringTransactionEncode(input_tmp, 3) + Transaction.Input /*string(input_tmp_str[:])*/
	//balance   := "sx"+StringTransactionEncode(Transaction.Balance/*[2:]*/, 32)

	proto, _ := json.Marshal(Transaction.Protocol)
	protocol := "tx" + string(proto)

	outResult := ""
	token := ""
	tokenBalance := ""
	vout := ""
	for i := 0; i < len(Transaction.Out); i++ {
		token_tmp := hex.EncodeToString([]byte(Transaction.Out[i].Token))
		token = StringTransactionEncode(string(token_tmp), 8)
		//balance_tmp   := strconv.FormatInt(int64(Transaction.Out[i].Balance.Uint64()), 16)
		tokenBalance = StringTransactionEncode(Transaction.Out[i].Balance, 32) //balance_tmp, 32)
		vout_tmp := strconv.Itoa(Transaction.Out[i].Vout)
		vout = StringTransactionEncode(vout_tmp, 3)
		outResult = outResult + "px" + vout + token + tokenBalance
	}
	sign := "mx" + StringTransactionEncode(Transaction.Sign, 128)
	publicKey := "nx" + StringTransactionEncode(Transaction.PublicKey, 64)
	tx := "rx" + StringTransactionEncode(Transaction.Tx, 64)
	balance := "sx" + StringTransactionEncode(Transaction.Balance /*[2:]*/, 32)
	//tx        := "gx"+StringTransactionEncode(Transaction.Tx, 64)

	re := from + to + nonce + fee + typ + outResult + input + balance + protocol + sign + publicKey + tx
	return re
}

//wite for change
func DecodeTransaction(transaction string) types.TransactionJson {
	g := strings.Index(transaction, "gx")
	h := strings.Index(transaction, "hx")
	i := strings.Index(transaction, "ix")
	k := strings.Index(transaction, "kx")
	l := strings.Index(transaction, "lx")
	m := strings.Index(transaction, "mx")
	n := strings.Index(transaction, "nx")
	s := strings.Index(transaction, "sx")
	To := /*"gx"+*/ transaction[g+4 : g+44]
	Nonce, _ := strconv.Atoi(transaction[h+2 : h+34])
	Fee := new(big.Int)
	Fee.SetString(transaction[i+2:i+34], 10)
	Type_tmp, _ := hex.DecodeString(transaction[k+2 : k+10])
	Type := string(Type_tmp)
	input_length, _ := strconv.ParseInt("0x"+transaction[l+2:l+5], 0, 64)
	Input_tmp, _ := hex.DecodeString(transaction[l+5 : l+5+int(input_length)])
	Input := string(Input_tmp)
	Sign := transaction[m+2 : m+130]
	PublicKey := transaction[n+2 : n+66]
	Balance := transaction[s+2 : s+34]
	trans := types.TransactionJson{
		//Out: outJson,
		Balance:   Balance,
		To:        To,
		Nonce:     Nonce,
		Fee:       (*Fee).String(),
		Type:      Type,
		Input:     Input,
		Sign:      Sign,
		PublicKey: PublicKey,
	}
	return trans
}

func StringTransactionEncode(feeString string, times int) string {
	feeStringLen := len(feeString)
	for i := 0; i < (times - feeStringLen); i++ {
		feeString = "0" + feeString
	}
	return feeString
}

func IntTransactionEncode(fee int, times int) string {
	feeString := strconv.Itoa(fee)
	feeStringLen := len(feeString)
	for i := 0; i < (times - feeStringLen); i++ {
		feeString = "0" + feeString
	}
	return feeString
}

func BigIntTransactionEncode(fee big.Int, times int) string {
	feeString := fee.String()
	feeStringLen := len(feeString)
	for i := 0; i < (times - feeStringLen); i++ {
		feeString = "0" + feeString
	}
	return feeString
}

// func TransactionProtocol(core_arg types.CoreStruct, Transaction types.TransactionJson) (bool, string) {
func TransactionProtocol(core_arg *types.CoreStruct, Transaction types.TransactionJson) (bool, string) {
	fmt.Println("$$$$$$$$$$$$$$$$$$$ [TransactionProtocol] $$$$$$$$$$$$$$$$$$$")
	fmt.Printf("%#v", Transaction)
	fmt.Println("\n[TransactionProtocol]Transaction.Type: ", Transaction.Type)
	fmt.Println("\n[TransactionProtocol]Transaction.Input: ", Transaction.Input)
	f := fmt.Sprintf("\n[TransactionProtocol]Transaction.From: [%s]", Transaction.From)
	t := fmt.Sprintf("\n[TransactionProtocol]Transaction.To: [%s]", Transaction.To)
	fmt.Println(f)
	fmt.Println(t)
	fmt.Println("\n[TransactionProtocol]Transaction.Nonce: ", Transaction.Nonce)
	fmt.Println("\n[TransactionProtocol]Transaction.PublicKey: ", Transaction.PublicKey)
	Transaction = TransactionModify("", Transaction)
	if Transaction.Type == "bnn" && Transaction.Input != "" {
		if Transaction.Input[0:6] == "0fa05f" && Transaction.Protocol != nil {
			//註冊
			//if(Transaction.Protocol!=nil){
			return CreateNewsStation(*core_arg, Transaction)
			//}
		}
		if Transaction.Input[0:6] == "293701" {
			//更新
			return UpdateNewsStation(*core_arg, Transaction)

		}
		if Transaction.Input[0:6] == "619759" {
			//刪除
		}
	}
	if Transaction.Type == "lock" {
		return LockProtocol(*core_arg, Transaction)
		//return result,"balance is not enough"
	}

	if Transaction.Type == "unlock" {
		return UnlockProtocol(*core_arg, Transaction)
		//return result,"balance is not enough"
	}

	if Transaction.Type == "vvmdeploy" {
		result, err := Deploy(Transaction.Input)
		if err != nil {
			return false, result
		}
		return true, result
	}

	if Transaction.Type == "vvmexcute" {
		result, err := Excute(Transaction.Input)
		if err != nil {
			return false, result
		}
		return true, result
	}
	// Create new coin
	if Transaction.Type == "90f4" {
		i, _ := strconv.Atoi(Transaction.Input[:1])
		fromAccountInfo := db.AccountHexGet(core_arg.Db, Transaction.From)
		inter := db.StringHexGet(core_arg.NameDb, Transaction.Input[1:1+i])

		if inter == "" {
			db.StringHexPut(core_arg.NameDb, Transaction.Input[1:1+i], Transaction.From)
			if fromAccountInfo.Token == nil {
				fromAccountInfo.Token = map[string]string{Transaction.Input[1 : 1+i]: Transaction.Input[1+i:]}
			} else {
				fromAccountInfo.Token[Transaction.Input[1:1+i]] = Transaction.Input[1+i:]
			}
			fmt.Println("[TransactionProtocol]fromAccountInfo.Token: ", fromAccountInfo)
			db.AccountHexPut(core_arg.Db, Transaction.From, fromAccountInfo)
		}
	}

	// Bolt Change
	if Transaction.Type == "bolt" {
		result, err := BoltChange(core_arg, Transaction)
		fmt.Println("[TransactionProtocol]err: ", err)
		return result, err
	}

	if len(Transaction.Input) > 6 {
		if Transaction.Input[0:6] == "82a353" {
			if len(Transaction.Input) > 20 {
				return false, "name too long"
			} else {
				var address string
				switch Transaction.Crypto {
				case "cic":
					address = crypto.CICKeyToAddress_hex(Transaction.PublicKey)
				case "secp256k1":
					address = crypto.CICKeyToAddress_hex(Transaction.PublicKey)
				case "eddsa":
					address = crypto.KeyToAddress_hex(Transaction.PublicKey)
				default:
					address = crypto.KeyToAddress_hex(Transaction.PublicKey)
				}
				fmt.Println("run", Transaction.Input[6:])
				x, y, _ := Name(*core_arg, address, Transaction.Input[6:])
				return x, y
			}
		}
	}

	fmt.Println("[TransactionProtocol] return] before VerifyTokenTransactionBalanceAndNonce")
	return VerifyTokenTransactionBalanceAndNonce(*core_arg, Transaction)
	// return false, error
	// return true, "success"
}

func UpdateTransactionDB(
	core_arg types.CoreStruct,
	fromAddress string,
	fromAccount types.AccountData,
	toAddress string,
	toAccount types.AccountData,
	feeAddress string,
	feeAccount types.AccountData) {
	db.AccountHexPut(core_arg.Db, feeAddress, feeAccount)
	db.AccountHexPut(core_arg.Db, fromAddress, fromAccount)
	db.AccountHexPut(core_arg.Db, toAddress, toAccount)
}

func UpdateTransactionDB4Bolt(
	core_arg types.CoreStruct,
	fromAddress string,
	fromAccount types.AccountData,
	toAddress string,
	toAccount types.AccountData,
	feeAddress string,
	feeAccount types.AccountData,
	ourAddress string,
	ourAccount types.AccountData) {
	db.AccountHexPut(core_arg.Db, feeAddress, feeAccount)
	db.AccountHexPut(core_arg.Db, fromAddress, fromAccount)
	db.AccountHexPut(core_arg.Db, toAddress, toAccount)
	db.AccountHexPut(core_arg.Db, ourAddress, ourAccount)
}

func UpdateFromTransactionDB(
	core_arg types.CoreStruct,
	fromAddress string,
	fromAccount types.AccountData,
	feeAddress string,
	feeAccount types.AccountData) {
	db.AccountHexPut(core_arg.Db, fromAddress, fromAccount)
	db.AccountHexPut(core_arg.Db, feeAddress, feeAccount)
}

func Transactiontest(SignBody types.SignBodys_Struct) types.TransactionJson {
	var trans types.TransactionJson
	Pri, _ := hex.DecodeString(SignBody.PrivateKey)
	codeInput := SignBody.Input
	var strpublicKey string
	var publicKey []byte
	switch SignBody.Crypto {
	case "eddsa":
		publicKey = eddsa.EddsaKeyToPublicKey(Pri)
	case "secp256k1":
		publicKey = secp.SecpKeyToPublicKey(Pri)
	case "cic":
		publicKey = secp.SecpKeyToPublicKey(Pri)
	case "a64":
		publicKey = secp.SecpKeyToPublicKey(Pri)
	default:
		publicKey = secp.SecpKeyToPublicKey(Pri)
	}
	strpublicKey = hex.EncodeToString(publicKey)

	if SignBody.Type[:3] == "vvm" {
		codeAddr := crypto.PubkeyToAddress(SignBody.Crypto, strpublicKey)
		codeInput = codeInput[:len(codeInput)-1] + ",\"account\":\"var sender = '" + codeAddr + "';\"" + codeInput[len(codeInput)-1:]
	}

	timetype := time.Unix(SignBody.Timelock.Time, 0)
	endtimetype := time.Unix(SignBody.Timelock.EndTime, 0)

	lock := types.TimeLock{
		Type:         SignBody.Timelock.Type,
		Amount:       SignBody.Timelock.Amount,
		Time:         timetype,
		EndTime:      endtimetype,
		UnlockAmount: SignBody.Timelock.UnlockAmount,
		IntervalTime: SignBody.Timelock.IntervalTime,
	}

	trans = types.TransactionJson{
		Balance:   SignBody.Balance,
		Out:       SignBody.Out,
		To:        SignBody.Address,
		Nonce:     SignBody.Nonce,
		Fee:       SignBody.Fee,
		Type:      SignBody.Type,
		Input:     codeInput,
		Crypto:    SignBody.Crypto,
		TimeLock:  lock,
		Protocol:  SignBody.Protocol,
		PublicKey: strpublicKey,
	}
	result := SignTransaction(Pri, trans)
	result = TransactionModify("", result)
	return result
}

func Transactiontestbatch(Nonce int, PrivateKey string) types.TransactionJson {
	var trans types.TransactionJson
	trans = types.TransactionJson{
		// Balance: "10",
		Balance: "0",
		To:      "7feca1b89235787bf69443e23467ec0e5a67bc74",
		Nonce:   Nonce,
		Type:    "bnn",
		Crypto:  "cic",
	}
	/*re := NewTransaction(
	            "10",
	            "7feca1b89235787bf69443e23467ec0e5a67bc74",
	            Nonce,
	            "10",//*big.NewInt(1),
	            "none",
				"none",
	    )*/
	a, _ := hex.DecodeString(PrivateKey)
	result := SignTransaction(a, trans)

	return result
}
