package transaction

import (
	"encoding/hex"
	"fmt"
	"lurcury/crypto"
	"lurcury/db"
	"lurcury/params"
	"lurcury/types"
	"math/big"
	"os"
	"strings"
)

func BoltChange(core_arg *types.CoreStruct, Transaction types.TransactionJson) (bool, string) {

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

	fromAccountInfo := db.AccountHexGet(core_arg.Db, address)
	// feeAccountInfo := db.AccountHexGet(core_arg.Db, params.Chain().Version.Eleve["dev"].FeeAddress)
	// toAccountInfo := db.AccountHexGet(core_arg.Db, Transaction.To)
	// ourAccountInfo := db.AccountHexGet(core_arg.Db, params.Chain().Version.Eleve["dev"].UsdtAddress)

	if Transaction.Nonce > fromAccountInfo.Nonce {
		return false, "nonce too high"
	}
	if Transaction.Nonce < fromAccountInfo.Nonce {
		return false, "nonce too low"
	}
	tokenName := strings.Split(Transaction.Input, "$")
	if len(tokenName) != 3 {
		return false, "params number must be 3"
	}
	exchangeRate := tokenName[2]

	// Bolt change balance
	tokenNameFrom := Transaction.Out[0].Token   // ex: usdtn
	changeBalance := Transaction.Out[0].Balance // ex: 10000000000(satoshi)  (100 usdtn)
	fmt.Println("[bolt] tokenNameFrom: ", tokenNameFrom)
	fmt.Println("[bolt] changeBalance: ", changeBalance)
	receiveBalance := new(big.Float)
	smallPartBalance := new(big.Float)
	changeBalance_float := new(big.Float)
	changeBalance_float.SetString(changeBalance)

	// Deal Transfer fee(TTN)
	// fromBalance := new(big.Int)
	// feeBalance := new(big.Int)
	// toBalance := new(big.Int)
	// fromMinusBalance := new(big.Int)
	// toAddBalance := new(big.Int)
	// transFeeBalance := new(big.Int)
	// receiveBalance := new(big.Float)
	// smallPartBalance := new(big.Float)
	// changeBalance_float := new(big.Float)

	// fromBalance.SetString(fromAccountInfo.Balance, 10)
	// feeBalance.SetString(feeAccountInfo.Balance, 10)
	// toBalance.SetString(toAccountInfo.Balance, 10)
	// // TODO
	// transFeeBalance = params.Chain().Version.Eleve["dev"].Fee
	// toAddBalance.SetString(Transaction.Balance, 10)
	// fromMinusBalance.Add(transFeeBalance, toAddBalance)

	// if fromBalance.Cmp(fromMinusBalance) >= 0 {
	// 	fromAccountInfo.Nonce = fromAccountInfo.Nonce + 1
	// 	fromAccountInfo.Balance = fromBalance.Sub(fromBalance, fromMinusBalance).String()
	// 	toAccountInfo.Balance = toBalance.Add(toBalance, toAddBalance).String()
	// 	feeAccountInfo.Balance = feeBalance.Add(feeBalance, transFeeBalance).String()
	// 	if tokenName[0] == tokenNameFrom { // "usdtn" == "usdtn"
	// 		// User Sub 100 USDTN
	// 		userTokenBalance1 := new(big.Float) // usdtn
	// 		changeBalance_float.SetString(changeBalance)
	// 		userTokenBalance1.SetString(fromAccountInfo.Token[tokenName[0]])
	// 		if userTokenBalance1.Cmp(changeBalance_float) >= 0 {
	// 			userTokenBalance1.Sub(userTokenBalance1, changeBalance_float)
	// 			fromAccountInfo.Token[tokenName[0]] = userTokenBalance1.Text('f', 0)
	// 			// 項目方 Add 99% USDTN
	// 			bigrate := new(big.Float)
	// 			bigrate.SetString("0.99") // 99%
	// 			bigPartBalance := new(big.Float)
	// 			bigPartBalance.Mul(changeBalance_float, bigrate)
	// 			companyTokenBalance1 := new(big.Float) // usdtn
	// 			companyTokenBalance1.SetString(toAccountInfo.Token[tokenName[0]])
	// 			companyTokenBalance1.Add(companyTokenBalance1, bigPartBalance)
	// 			toAccountInfo.Token[tokenName[0]] = companyTokenBalance1.Text('f', 0)
	// 		} else {
	// 			return false, tokenName[0] + " balance not enough"
	// 		}
	// 	}
	// } else {
	// 	return false, "balance not enough"
	// }
	// fmt.Println("-------------- Final All Account content: ----------------")
	// fmt.Println("fromAccountInfo: ", fromAccountInfo)
	// fmt.Println("toAccountInfo: ", toAccountInfo)
	// fmt.Println("ourAccountInfo: ", ourAccountInfo)

	// Tx: User -> MCC 項目方 (100 usdtn)
	// var u types.SignBodys_Struct
	// u.Type = "ttn"
	// u.Fee = params.Chain().Version.Eleve["dev"].Fee.String()
	// u.Address = Transaction.To
	// u.Out = []types.TransactionOut{
	// 	types.TransactionOut{
	// 		Token:   tokenName[0],
	// 		Balance: Transaction.Out[0].Balance,
	// 	}}
	// u.Balance = "0"
	// u.Nonce = db.AccountHexGet(core_arg.Db, Transaction.From).Nonce
	// u.Input = ""
	// // TODO!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	// // u.PrivateKey = "ae1c0127e663c7048e42b61a5c99f7ee125c19ee129b8ddea81b5090127f7822"
	// u.Crypto = "cic"
	// u.Protocol = nil
	// firstTransaction := Transactiontest(u)
	// res1, err := VerifyTokenTransactionBalanceAndNonce(*core_arg, firstTransaction)
	// if res1 == false {
	// 	fmt.Println("res1 == false")
	// 	return false, err
	// }

	Transaction.Type = "ttn"
	Transaction.Input = "go"
	fmt.Println("@@@@@@@@@@ Transaction.Nonce: ", Transaction.Nonce)
	fmt.Println("@@@@@@@@@@ Transaction.From.Nonce: ", db.AccountHexGet(core_arg.Db, Transaction.From).Nonce)
	Transaction.Nonce = db.AccountHexGet(core_arg.Db, Transaction.From).Nonce
	// res1, err := VerifyTokenTransactionBalanceAndNonce(*core_arg, Transaction)
	// if res1 == false {
	// 	fmt.Println("res1 == false")
	// 	return false, err
	// }

	// Tx: MCC項目方 Sub 500 MCC, Add 500 MCC to User
	var s types.SignBodys_Struct
	const prec = 200
	receiveBalance.SetPrec(prec)
	rate := new(big.Float)
	ethWei := new(big.Float)
	ethWei.SetString("10000000000") // 10^10 (omni[10^8] -> ethwei[10^18] )
	rate.SetString(exchangeRate)    // ex: 5
	receiveBalance.Mul(changeBalance_float, rate)
	receiveBalance.Mul(receiveBalance, ethWei)
	s.Type = "ttn"
	s.Fee = params.Chain().Version.Eleve["dev"].Fee.String()
	s.Address = Transaction.From
	s.Out = []types.TransactionOut{
		types.TransactionOut{
			Token:   tokenName[1], // MCC or TTN
			Balance: receiveBalance.Text('f', 0),
		}}
	s.Balance = "0"
	s.Nonce = db.AccountHexGet(core_arg.Db, Transaction.To).Nonce
	s.Input = ""
	// 項目方PrivateKey
	// s1.PrivateKey = os.Getenv("MCC" + "PRI")
	s.PrivateKey = os.Getenv(strings.ToUpper(tokenName[1]) + "PRI") // "MCC" + "PRI" or "TTN"+"PRI"
	s.Crypto = "cic"
	s.Protocol = nil
	secondTransaction := Transactiontest(s)
	// res2, err := VerifyTokenTransactionBalanceAndNonce(*core_arg, secondTransaction)
	// if res2 == false {
	// 	fmt.Println("res2 == false")
	// 	return false, err
	// }

	// Tx: MCC項目方Usdt 1% -> Our Usdt address
	var t types.SignBodys_Struct
	smallrate := new(big.Float)
	smallrate.SetString("0.01") // 1%
	smallPartBalance.Mul(changeBalance_float, smallrate)
	t.Type = "ttn"
	t.Fee = params.Chain().Version.Eleve["dev"].Fee.String()
	t.Address = params.Chain().Version.Eleve["dev"].UsdtAddress
	t.Out = []types.TransactionOut{
		types.TransactionOut{
			Token:   tokenName[0],
			Balance: smallPartBalance.Text('f', 0),
		}}
	t.Balance = "0"
	t.Nonce = db.AccountHexGet(core_arg.Db, Transaction.To).Nonce + 1 // Continue after secondtransaction, so need to add 1
	t.Input = ""
	// 項目方PrivateKey
	t.PrivateKey = os.Getenv(strings.ToUpper(tokenName[1]) + "PRI") // MCCPRI or TTNPRI
	t.Crypto = "cic"
	t.Protocol = nil
	thirdTransaction := Transactiontest(t)
	// res3, err := VerifyTokenTransactionBalanceAndNonce(*core_arg, thirdTransaction)
	// if res3 == false {
	// 	fmt.Println("res3 == false")
	// 	return false, err
	// }

	// Exchange: ===============   USDT -> TTN   =============
	if tokenName[1] == "ttn" {
		// User 100 usdt -> Our usdt address
		result, err := VerifyTokenTransactionBalanceAndNonce(*core_arg, Transaction)
		if result == true {

			// Our usdt address 500 TTN -> User address
			var s1 types.SignBodys_Struct
			const prec = 200
			receiveBalance.SetPrec(prec)
			rate := new(big.Float)
			ethWei := new(big.Float)
			ethWei.SetString("10000000000") // 10^10 (omni[10^8] -> ethwei[10^18] )
			rate.SetString(exchangeRate)    // ex: 5
			receiveBalance.Mul(changeBalance_float, rate)
			receiveBalance.Mul(receiveBalance, ethWei)

			s1.Type = "ttn"
			s1.Fee = params.Chain().Version.Eleve["dev"].Fee.String()
			s1.Address = Transaction.From
			s1.Out = nil
			s1.Balance = receiveBalance.Text('f', 0)
			s1.Nonce = db.AccountHexGet(core_arg.Db, Transaction.To).Nonce
			s1.Input = ""
			s1.PrivateKey = os.Getenv(strings.ToUpper(tokenName[0]) + "PRI") // "USDTN"+ "PRI"
			s1.Crypto = "cic"
			s1.Protocol = nil
			ttnTransaction := Transactiontest(s1)

			ttnTransaction.Tx = hex.EncodeToString(crypto.Keccak256([]byte(EncodeForSign(secondTransaction))))
			db.TransactionHexPut(core_arg.Db, ttnTransaction.Tx, ttnTransaction)
			core_arg.PendingTransaction = append(core_arg.PendingTransaction, ttnTransaction)

			return true, "success"
		} else {

			fmt.Println("[BoltChange token = ttn]: false, err")
			return false, err
		}
	}

	// result := true
	// if core_arg.Model == "1" || core_arg.Model == "2" {
	// 	toAccountInfo.Transaction = append(toAccountInfo.Transaction, Transaction)
	// 	fromAccountInfo.Transaction = append(fromAccountInfo.Transaction, Transaction)

	// 	// // Tx: 項目方MCC -> User MCC
	// 	// toAccountInfo.Transaction = append(toAccountInfo.Transaction, secondTransaction)
	// 	// fromAccountInfo.Transaction = append(fromAccountInfo.Transaction, secondTransaction)

	// 	// // Tx: 項目方Usdt 1% -> Our Usdt address
	// 	// toAccountInfo.Transaction = append(toAccountInfo.Transaction, thirdTransaction)
	// 	// ourAccountInfo.Transaction = append(ourAccountInfo.Transaction, thirdTransaction)

	// }

	// TODO
	// if core_arg.Model == "2" {
	// 	if len(toAccountInfo.Transaction) > 10 {
	// 		for i := 0; i < 10; i++ {
	// 			toAccountInfo.Transaction[i] = toAccountInfo.Transaction[i+1]
	// 		}
	// 		toAccountInfo.Transaction = toAccountInfo.Transaction[:9]
	// 	}
	// 	if len(fromAccountInfo.Transaction) > 10 {
	// 		for i := 0; i < 10; i++ {
	// 			fromAccountInfo.Transaction[i] = fromAccountInfo.Transaction[i+1]
	// 		}
	// 		fromAccountInfo.Transaction = fromAccountInfo.Transaction[:9]
	// 	}
	// }

	// Transaction.Tx = hex.EncodeToString(crypto.Keccak256([]byte(EncodeForSign(Transaction))))
	// Transaction = TransactionModify(Transaction.BlockHash, Transaction)

	// secondTransaction.Tx = hex.EncodeToString(crypto.Keccak256([]byte(EncodeForSign(secondTransaction))))
	// thirdTransaction.Tx = hex.EncodeToString(crypto.Keccak256([]byte(EncodeForSign(thirdTransaction))))
	// fmt.Println("1. Transaction.Tx:::", Transaction.Tx)
	// fmt.Println("2. secondTransaction.Tx:::", secondTransaction.Tx)
	// fmt.Println("3. thirdTransaction.Tx:::", thirdTransaction.Tx)

	// if result == true {
	// if true {
	// 	UpdateTransactionDB4Bolt(
	// 		*core_arg,
	// 		address, fromAccountInfo,
	// 		params.Chain().Version.Eleve["dev"].FeeAddress, feeAccountInfo,
	// 		Transaction.To, toAccountInfo,
	// 		params.Chain().Version.Eleve["dev"].UsdtAddress, ourAccountInfo)
	// 	// TODO
	// 	// UpdateTransactionDB(
	// 	// 	*core_arg,
	// 	// 	address, fromAccountInfo,
	// 	// 	params.Chain().Version.Eleve["dev"].FeeAddress, feeAccountInfo,
	// 	// 	Transaction.To, toAccountInfo)
	// } else {
	// 	return false, "[bolt] token not enough"
	// }
	// db.TransactionHexPut(core_arg.Db, Transaction.Tx, Transaction)
	// db.TransactionHexPut(core_arg.Db, secondTransaction.Tx, secondTransaction)
	// db.TransactionHexPut(core_arg.Db, thirdTransaction.Tx, thirdTransaction)

	// core_arg.PendingTransaction = append(core_arg.PendingTransaction, firstTransaction)
	core_arg.PendingTransaction = append(core_arg.PendingTransaction, Transaction)
	core_arg.PendingTransaction = append(core_arg.PendingTransaction, secondTransaction)
	core_arg.PendingTransaction = append(core_arg.PendingTransaction, thirdTransaction)

	return true, "success"
}
