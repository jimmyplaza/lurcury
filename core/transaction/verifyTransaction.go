package transaction

import (
	"encoding/hex"
	"fmt"
	"lurcury/crypto"
	eddsa "lurcury/crypto/eddsa"
	secp "lurcury/crypto/secp256k1"
	"lurcury/db"
	"lurcury/params"
	"lurcury/types"
	"math/big"
	//"time"
)

func VerifyTransactionSign(Transaction types.TransactionJson) bool {
	pub, _ := hex.DecodeString(Transaction.PublicKey)
	// fmt.Println("[VerifyTransactionSign] Transaction.PublicKey: ", Transaction.PublicKey)
	// fmt.Printf("%#v ", Transaction)
	msg := crypto.Keccak256([]byte(EncodeForSign(Transaction)))
	if Transaction.Crypto == "a64" {
		fmt.Println("Transaction.Crypto is a64!!")
		msg = crypto.Sha256([]byte(EncodeForSign_Simple(Transaction)))
	}
	sign, _ := hex.DecodeString(Transaction.Sign)
	// fmt.Println("[VerifyTransactionSign] hex.EncodeToString(msg): ", hex.EncodeToString(msg))
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
func VerifyHttpTransactionBalanceAndNonce(core_arg types.CoreStruct, Transaction types.TransactionJson) (bool, string, string) {
	//address := crypto.KeyToAddress_hex(Transaction.PublicKey)
	var address string
	fmt.Println("[VerifyHttpTransactionBalanceAndNonce]********: Transaction.Crypto: ", Transaction.Crypto)
	switch Transaction.Crypto {
	case "cic":
		address = crypto.CICKeyToAddress_hex(Transaction.PublicKey)
	case "secp256k1":
		address = crypto.CICKeyToAddress_hex(Transaction.PublicKey)
	case "eddsa":
		address = crypto.KeyToAddress_hex(Transaction.PublicKey)
	default:
		fmt.Println("******* this is default Transaction.Crypto")
		address = crypto.CICKeyToAddress_hex(Transaction.PublicKey)
	}
	if address == Transaction.To {
		return false, "cannot send to urself", ""
	}
	fmt.Println("[VerifyHttpTransactionBalanceAndNonce]********: core_arg db: ", core_arg.Db)
	fmt.Println("[VerifyHttpTransactionBalanceAndNonce]********: Transaction.PublicKey: ", Transaction.PublicKey)
	fmt.Println("[VerifyHttpTransactionBalanceAndNonce]********: address: ", address)
	fromAccountInfo := db.AccountHexGet(core_arg.Db, address)
	fmt.Println("[VerifyHttpTransactionBalanceAndNonce]********: fromAccountInfo balance: ", fromAccountInfo.Balance)
	fmt.Println("[VerifyHttpTransactionBalanceAndNonce]********: fromAccountInfo transaction: ", fromAccountInfo.Transaction)
	fmt.Println("[VerifyHttpTransactionBalanceAndNonce]********: fromAccountInfo nonce: ", fromAccountInfo.Nonce)
	fmt.Println("[VerifyHttpTransactionBalanceAndNonce]********: fromAccountInfo token: ", fromAccountInfo.Token)

	if Transaction.Nonce < fromAccountInfo.Nonce {
		return false, "nonce too low", ""
	}
	fromBalance := new(big.Int)
	fromMinusBalance := new(big.Int)
	toAddBalance := new(big.Int)
	transFeeBalance := new(big.Int)

	fromBalance.SetString(fromAccountInfo.Balance, 10)
	transFeeBalance = params.Chain().Version.Eleve["dev"].Fee // 0.1 TTN

	toAddBalance.SetString(Transaction.Balance, 10)
	fromMinusBalance.Add(transFeeBalance, toAddBalance) // 0.1 TTN + transfer balance
	fmt.Println("[VerifyHttpTransactionBalanceAndNonce]********: toAddBalance: ", toAddBalance)
	fmt.Println("[VerifyHttpTransactionBalanceAndNonce]********: fromBalance: ", fromBalance)
	fmt.Println("[VerifyHttpTransactionBalanceAndNonce]********: fromMinusBalance: ", fromMinusBalance)
	if VerifyTransactionSign(Transaction) == false {
		return false, "error sign", ""
	}

	fmt.Println("[VerifyHttpTransactionBalanceAndNonce]sign:", crypto.Keccak256([]byte(EncodeForSign(Transaction))))
	if fromBalance.Cmp(fromMinusBalance) >= 0 {
		//return true, "success", hex.EncodeToString(crypto.Keccak256([]byte(EncodeForSign(Transaction))))
		//EncodeString := hex.EncodeToString(crypto.Keccak256([]byte(EncodeForSign(Transaction))))
		Transaction = TransactionModify("", Transaction)
		//if(Transaction.Crypto=="a64"){
		//	EncodeString = hex.EncodeToString(crypto.Sha256([]byte(EncodeForSign_Simple(Transaction))))
		//}
		return true, "success", Transaction.Tx //EncodeString
	}
	fmt.Println("[VerifyHttpTransactionBalanceAndNonce] return")
	return false, "balance not enough", ""
}

func VerifyTokenTransactionBalanceAndNonce(core_arg types.CoreStruct, Transaction types.TransactionJson) (bool, string) {
	fmt.Println("\n **************** [VerifyTokenTransactionBalanceAndNonce]")
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

	if len(Transaction.To) < 30 {
		inter := db.StringHexGet(core_arg.NameDb, Transaction.To) //, &inter)
		Transaction.To = string(inter)
	}

	//address := crypto.KeyToAddress_hex(Transaction.PublicKey)
	fromAccountInfo := db.AccountHexGet(core_arg.Db, address)
	feeAccountInfo := db.AccountHexGet(core_arg.Db, params.Chain().Version.Eleve["dev"].FeeAddress)
	toAccountInfo := db.AccountHexGet(core_arg.Db, Transaction.To)
	fmt.Println("[VerifyTokenTransactionBalanceAndNonce]############## 1. Ready db.AccountHexGet, address: ", address)
	fmt.Println("[VerifyTokenTransactionBalanceAndNonce]############## 2. Get FeeAddress: ", params.Chain().Version.Eleve["dev"].FeeAddress)
	fmt.Println("[VerifyTokenTransactionBalanceAndNonce]############## 3. Get Transaction.To address: ", Transaction.To)
	// When toAccountInfo is null
	fmt.Println("[VerifyTokenTransactionBalanceAndNonce]-----------------: toAccountInfo: ", toAccountInfo)
	// if reflect.DeepEqual(toAccountInfo, types.AccountData{}) {
	// 	return false, "toAccount format is wrong, got null"
	// }
	fmt.Println("[VerifyTokenTransactionBalanceAndNonce]############## 4. this Transaction.Nonce: ", Transaction.Nonce)
	fmt.Println("[VerifyTokenTransactionBalanceAndNonce]############## 5. fromAccountInfo.Nonce: ", fromAccountInfo.Nonce)

	if Transaction.Nonce > fromAccountInfo.Nonce {
		fmt.Println("[VerifyTokenTransactionBalanceAndNonce] nonce too high")
		return false, "nonce too high"
	}
	if Transaction.Nonce < fromAccountInfo.Nonce {
		fmt.Println("[VerifyTokenTransactionBalanceAndNonce] nonce too low")
		return false, "nonce too low"
	}
	fmt.Println("[VerifyTokenTransactionBalanceAndNonce]############## fromAccountInfo.Balance: ", fromAccountInfo.Balance)
	fmt.Println("[VerifyTokenTransactionBalanceAndNonce]############## toAccountInfo.Balance: ", toAccountInfo.Balance)
	fromBalance := new(big.Int)
	feeBalance := new(big.Int)
	toBalance := new(big.Int)
	fromMinusBalance := new(big.Int)
	toAddBalance := new(big.Int)
	transFeeBalance := new(big.Int)
	fromBalance.SetString(fromAccountInfo.Balance, 10)
	feeBalance.SetString(feeAccountInfo.Balance, 10)
	toBalance.SetString(toAccountInfo.Balance, 10)
	//transFeeBalance.SetString(/*Transaction.Fee*/params.Chain().Version.Eleve["dev"].Fee,10)
	transFeeBalance = params.Chain().Version.Eleve["dev"].Fee
	toAddBalance.SetString(Transaction.Balance, 10)
	fromMinusBalance.Add(transFeeBalance, toAddBalance)
	//fmt.Println("fromBalance.Cmp",fromBalance)
	//fmt.Println("fromMinusBalance",fromMinusBalance)

	if fromBalance.Cmp(fromMinusBalance) >= 0 {
		fromAccountInfo.Nonce = fromAccountInfo.Nonce + 1
		fromAccountInfo.Balance = fromBalance.Sub(fromBalance, fromMinusBalance).String()
		toAccountInfo.Balance = toBalance.Add(toBalance, toAddBalance).String()
		feeAccountInfo.Balance = feeBalance.Add(feeBalance, transFeeBalance).String()
		//fromAccountInfo.Transaction= append(fromAccountInfo.Transaction, Transaction)
		//toAccountInfo.Transaction= append(toAccountInfo.Transaction, Transaction)
	} else {
		return false, "[VerifyTokenTransactionBalanceAndNonce] balance not enough"
	}
	result, fromAccountInfo, toAccountInfo := VerifyTokenBalance(Transaction, fromAccountInfo, toAccountInfo)
	if core_arg.Model == "1" || core_arg.Model == "2" {
		toAccountInfo.Transaction = append(toAccountInfo.Transaction, Transaction)
		fromAccountInfo.Transaction = append(fromAccountInfo.Transaction, Transaction)
	}
	if core_arg.Model == "2" {
		if len(toAccountInfo.Transaction) > 10 {
			for i := 0; i < 10; i++ {
				toAccountInfo.Transaction[i] = toAccountInfo.Transaction[i+1]
			}
			toAccountInfo.Transaction = toAccountInfo.Transaction[:9]
		}
		if len(fromAccountInfo.Transaction) > 10 {
			for i := 0; i < 10; i++ {
				fromAccountInfo.Transaction[i] = fromAccountInfo.Transaction[i+1]
			}
			fromAccountInfo.Transaction = fromAccountInfo.Transaction[:9]
		}
	}

	Transaction.Tx = hex.EncodeToString(crypto.Keccak256([]byte(EncodeForSign(Transaction))))
	Transaction = TransactionModify(Transaction.BlockHash, Transaction)
	// fmt.Println("Transaction.Tx:::", Transaction.Tx)
	//if(Transaction.Crypto=="a64"){
	//        Transaction.Tx = hex.EncodeToString(crypto.Sha256([]byte(EncodeForSign_Simple(Transaction))))
	//}
	if result == true {
		UpdateTransactionDB(core_arg,
			address, fromAccountInfo,
			params.Chain().Version.Eleve["dev"].FeeAddress, feeAccountInfo,
			Transaction.To, toAccountInfo)
	} else {
		return false, "[VerifyTokenTransactionBalanceAndNonce] token not enough"
	}
	db.TransactionHexPut(core_arg.Db, Transaction.Tx, Transaction)
	return true, "success"
}

func VerifyFee(transaction types.TransactionJson, fromAccount types.AccountData, feeAccount types.AccountData) (bool, types.AccountData, types.AccountData) {
	fromBalance := new(big.Int)
	feeBalance := new(big.Int)
	transBalance := new(big.Int)
	fromBalance.SetString(fromAccount.Balance, 10)
	feeBalance.SetString(feeAccount.Balance, 10)
	transBalance.SetString(transaction.Fee, 10)
	fmt.Println("fromAccount.Balance:", fromAccount.Balance, "feeAccount.Balance:", feeAccount.Balance, "transaction.Balance", transaction.Fee)
	if fromBalance.Cmp(transBalance) >= 0 {
		fromBalance.Sub(fromBalance, transBalance)
		feeBalance.Add(feeBalance, transBalance)
		fromAccount.Balance = fromBalance.String()
		feeAccount.Balance = feeBalance.String()
		return true, fromAccount, feeAccount
	} else {
		return false, fromAccount, feeAccount
	}
}

func VerifyBalance(transaction types.TransactionJson, fromAccount types.AccountData, toAccount types.AccountData) (bool, types.AccountData, types.AccountData) {
	fromBalance := new(big.Int)
	toBalance := new(big.Int)
	transBalance := new(big.Int)
	fromBalance.SetString(fromAccount.Balance, 10)
	toBalance.SetString(toAccount.Balance, 10)
	transBalance.SetString(transaction.Balance, 10)
	fmt.Println("fromAccount.Balance:", fromAccount.Balance, "toAccount.Balance:", toAccount.Balance, "transaction.Balance:", transaction.Balance)
	fmt.Println("fromBalance.Cmp(transBalance):", fromBalance.Cmp(transBalance), "fromBalance:", fromBalance, "transBalance:", transBalance)
	if fromBalance.Cmp(transBalance) >= 0 {
		fromBalance.Sub(fromBalance, transBalance)
		toBalance.Add(toBalance, transBalance)
		fromAccount.Balance = fromBalance.String()
		toAccount.Balance = toBalance.String()
		return true, fromAccount, toAccount
	} else {
		return false, fromAccount, toAccount
	}
}
func VerifyTokenBalance(transaction types.TransactionJson, fromAccount types.AccountData, toAccount types.AccountData) (bool, types.AccountData, types.AccountData) {
	fromBalance := new(big.Int)
	toBalance := new(big.Int)
	transBalance := new(big.Int)
	for i := 0; i < len(transaction.Out); i++ {
		tokenName := transaction.Out[i].Token
		fromBalance.SetString(fromAccount.Token[tokenName], 10)
		toBalance.SetString(toAccount.Token[tokenName], 10)
		transBalance.SetString(transaction.Out[i].Balance, 10)
		if fromBalance.Cmp(transBalance) >= 0 {
			fromBalance.Sub(fromBalance, transBalance)
			toBalance.Add(toBalance, transBalance)
			fromAccount.Token[tokenName] = fromBalance.String()
			if toAccount.Token == nil {
				toAccount.Token = map[string]string{tokenName: toBalance.String()}
			} else {
				toAccount.Token[tokenName] = toBalance.String()
			}
		} else {
			return false, fromAccount, toAccount
		}
	}
	return true, fromAccount, toAccount
}
