package account

import (
	//"encoding/hex"
	"fmt"
	"lurcury/db"
	"lurcury/types"
	"testing"
)

func TestAccount(t *testing.T){
	a := InitAccount("8634d39c4affcfc7324427bfb331ee588e6e09cb", "club", "1000000000000000000000000000")
	fmt.Println(a)
	core_arg := &types.CoreStruct{}
	core_arg.Db = db.OpenDB("./123")
	db.AccountHexPut(core_arg.Db, "8634d39c4affcfc7324427bfb331ee588e6e09cb", a)
	/*
	check := func(f string, got, want interface{}) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("%s mismatch: got %v, want %v", f, got, want)
		}
	}

        data := []types.BalanceData{{Token:"abc", Balance: *big.NewInt(1)}}
        data = append(data,types.BalanceData{Token:"abc", Balance: *big.NewInt(1)})
        accountdata := types.AccountData{Address:"123",Nonce:1,Balance:data}
	a,b,c := NewAccount()
	fmt.Println("privateKey:", a,"publicKey:", b,"address:", c)
	check("go",accountdata.Address,"123")
	*/
}
