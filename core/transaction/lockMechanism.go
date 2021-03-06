package transaction

import (
	"fmt"
	"lurcury/crypto"
	"lurcury/db"
	"lurcury/params"
        "lurcury/types"
	"math/big"
	"time"
)

func DeletUnlockRecord(core_arg *types.AccountData, i int)(*types.AccountData){
        core_arg.TimeLock[i] = core_arg.TimeLock[len(core_arg.TimeLock)-1]
        core_arg.TimeLock = core_arg.TimeLock[:len(core_arg.TimeLock)-1]
        return core_arg
}

func LockProtocol(core_arg types.CoreStruct, trans types.TransactionJson)(bool, string){
        var address string
        switch trans.Crypto {
                case "cic":
                        address = crypto.CICKeyToAddress_hex(trans.PublicKey)
                case "secp256k1":
                        address = crypto.CICKeyToAddress_hex(trans.PublicKey)
                case "eddsa":
                        address = crypto.KeyToAddress_hex(trans.PublicKey)
                default:
                        address = crypto.KeyToAddress_hex(trans.PublicKey)
        }

	fromAccountInfo := db.AccountHexGet(core_arg.Db, address)
	toAccountInfo := db.AccountHexGet(core_arg.Db, trans.To)
        fromBalance := new(big.Int)
        fromBalance.SetString(fromAccountInfo.Balance, 10)
        transBalance := new(big.Int)
        transBalance.SetString(trans.TimeLock.Amount, 10)

        feeAccountInfo := db.AccountHexGet(core_arg.Db, params.Chain().Version.Eleve["dev"].FeeAddress)
        feeBalance := new(big.Int)
	fromMinusBalance := new(big.Int)
        transFeeBalance := new(big.Int)
        feeBalance.SetString(feeAccountInfo.Balance,10)
        transFeeBalance = params.Chain().Version.Eleve["dev"].Fee
	fromMinusBalance.Add(transFeeBalance, transBalance)

	if(fromBalance.Cmp(fromMinusBalance)>=0){
		toAccountInfo.TimeLock = append(toAccountInfo.TimeLock, trans.TimeLock)
		fromAccountInfo.Balance = fromBalance.Sub(fromBalance, fromMinusBalance).String()
		feeAccountInfo.Balance = feeBalance.Add(feeBalance, transFeeBalance).String()
		fromAccountInfo.Balance = fromBalance.String()
		fromAccountInfo.Nonce = fromAccountInfo.Nonce + 1

		UpdateTransactionDB(core_arg,
			address, 
			fromAccountInfo,
			params.Chain().Version.Eleve["dev"].FeeAddress, 
			feeAccountInfo,
			trans.To, 
			toAccountInfo,
		)

		return true, "success"
	}
	return false, "balance is not enough"
}

//??????????????????
func UnlockProtocol(core_arg types.CoreStruct, trans types.TransactionJson)(bool, string){
        var address string
        switch trans.Crypto {
                case "cic":
                        address = crypto.CICKeyToAddress_hex(trans.PublicKey)
                case "secp256k1":
                        address = crypto.CICKeyToAddress_hex(trans.PublicKey)
                case "eddsa":
                        address = crypto.KeyToAddress_hex(trans.PublicKey)
                default:
                        address = crypto.KeyToAddress_hex(trans.PublicKey)
        }

	fromAccountInfo := db.AccountHexGet(core_arg.Db, address)
        fromBalance := new(big.Int)
        fromBalance.SetString(fromAccountInfo.Balance, 10)
        feeAccountInfo := db.AccountHexGet(core_arg.Db, params.Chain().Version.Eleve["dev"].FeeAddress)
        feeBalance := new(big.Int)
        transFeeBalance := new(big.Int)
        feeBalance.SetString(feeAccountInfo.Balance,10)
        transFeeBalance = params.Chain().Version.Eleve["dev"].Fee
	if(fromBalance.Cmp(transFeeBalance)<0){
		return false, "balance not enough"
	}
	fromAccountInfo.Balance = fromBalance.Sub(fromBalance, transFeeBalance).String()

	now := time.Unix(trans.Timestamp, 0)
	for index, element:= range fromAccountInfo.TimeLock{
		if element.Type == "linearlock"{
			if(now.After(element.Time)){
				interV := 0.1
				if(now.After(element.EndTime)){
					interV = 1
				}else{
					interV = (float64(now.Unix())-float64(element.Time.Unix()))/(float64(element.EndTime.Unix())-float64(element.Time.Unix()))
				}
 				fmt.Println("?????????????????????:", interV)
                                ratev := big.NewFloat(interV)
                                lockAmount := new(big.Float)
                                lockAmount.SetString(element.Amount)
                                fmt.Println("?????????:", lockAmount)
                                unlockAmount := new(big.Float)
                                unlockAmount.SetString(element.UnlockAmount)
                                fmt.Println("???????????????:", unlockAmount)
                                ratev.Mul(ratev,lockAmount)
                                fmt.Println("???????????????:", ratev)
                                ratev.Sub(ratev,unlockAmount)
                                result := new(big.Int)
                                ratev.Int(result)
                                fmt.Println("???????????????:",result)
                                unlockAmounts := new(big.Int)
                                unlockAmount.Int(unlockAmounts)
                                unlockAmounts.Add(unlockAmounts, result)
                                fmt.Println("???????????????:", unlockAmounts)
                                fromBalance.Add(fromBalance, result)
                                fromAccountInfo.Balance = fromBalance.String()
                                fmt.Println("????????????:", fromAccountInfo.Balance)
                                fromAccountInfo.TimeLock[index].UnlockAmount = unlockAmounts.String()
                                //fromAccountInfo.Nonce = fromAccountInfo.Nonce + 1
                                //if(now.After(element.EndTime)){
                                //       DeletUnlockRecord(&fromAccountInfo, index)
                                //}
			}
                        //?????????????????????
			//???????????????
			//??????????????????
			//??????????????????
			//????????????????????????
		}
		if element.Type == "pointlock"{
			if(now.After(element.EndTime)){
				lockAmount := new(big.Int)
				lockAmount.SetString(element.Amount, 10)
				fromBalance.Add(fromBalance, lockAmount)
				fromAccountInfo.Balance = fromBalance.String()
				//fromAccountInfo.Nonce = fromAccountInfo.Nonce + 1
				//DeletUnlockRecord(&fromAccountInfo, index)
			}
		}
                if element.Type == "ladderlock"{
                        if(now.After(element.Time)){
                                interV := 0.1
                                if(now.After(element.EndTime)){
                                        interV = 1
                                }else{
                                        interV = float64(int((float64(now.Unix())-float64(element.Time.Unix()))/float64(element.IntervalTime/*.Unix()*/)))
					
					fmt.Println("????????????:", interV)
					total := (float64(element.EndTime.Unix())-float64(element.Time.Unix()))/float64(element.IntervalTime/*.Unix()*/)
					interV = interV/total
                                }
 				fmt.Println("?????????????????????:", interV)
                                ratev := big.NewFloat(interV)
                                lockAmount := new(big.Float)
                                lockAmount.SetString(element.Amount)
                                fmt.Println("?????????:", lockAmount)
                                unlockAmount := new(big.Float)
                                unlockAmount.SetString(element.UnlockAmount)
                                fmt.Println("???????????????:", unlockAmount)
                                ratev.Mul(ratev,lockAmount)
                                fmt.Println("???????????????:", ratev)
                                ratev.Sub(ratev,unlockAmount)
                                result := new(big.Int)
                                ratev.Int(result)
                                fmt.Println("???????????????:",result)
                                unlockAmounts := new(big.Int)
                                unlockAmount.Int(unlockAmounts)
                                unlockAmounts.Add(unlockAmounts, result)
                                fmt.Println("???????????????:", unlockAmounts)
                                fromBalance.Add(fromBalance, result)
                                fromAccountInfo.Balance = fromBalance.String()
                                fmt.Println("????????????:", fromAccountInfo.Balance)
                                fromAccountInfo.TimeLock[index].UnlockAmount = unlockAmounts.String()
                                //fromAccountInfo.Nonce = fromAccountInfo.Nonce + 1
                                //if(now.After(element.EndTime)){
                                //        DeletUnlockRecord(&fromAccountInfo, index)
                                //}
                        }
                        //?????????????????????
                        //???????????????
                        //??????????????????
                        //??????????????????
                        //????????????????????????
                }
	}
	fromAccountInfo.Nonce = fromAccountInfo.Nonce + 1
	s := 0
	for _, element:= range fromAccountInfo.TimeLock{
                if(now.After(element.EndTime)){
                	DeletUnlockRecord(&fromAccountInfo, s)
		}else{
			s =s +1
		}
	}

        UpdateFromTransactionDB(core_arg,
                        address, 
                        fromAccountInfo,
                        params.Chain().Version.Eleve["dev"].FeeAddress, 
                        feeAccountInfo,
        )


	db.AccountHexPut(core_arg.Db, address, fromAccountInfo)
	return true, "success"
}


