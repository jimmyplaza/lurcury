package transaction

import (
        //"encoding/hex"
        "fmt"
	"lurcury/account"
	//"lurcury/crypto"
	"lurcury/db"
        "lurcury/types"
        //"math/big"
        //"reflect"
        "testing"
)

func TestTransaction(t *testing.T){
        core_arg := &types.CoreStruct{}
        core_arg.Db = db.OpenDB("../dbdata")
        account_tmp := account.Account_exp()
        db.AccountHexPut(core_arg.Db, account_tmp.Address, account_tmp)
        fmt.Println(db.AccountHexGet(core_arg.Db, "gx5ee464a101d58877f00957eff452c148e7f75833"))
        account_tmp.Address = "gx"
        db.AccountHexPut(core_arg.Db, account_tmp.Address, account_tmp)
        fmt.Println(db.AccountHexGet(core_arg.Db, "gx5ee464a101d58877f00957eff452c148e7f75833"))

}
