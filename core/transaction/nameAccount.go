package transaction

import(
        "fmt"
        "math/big"
        "lurcury/db"
        "lurcury/params"
        "lurcury/types"
)

func Name(core_arg types.CoreStruct ,address string, name string)(bool, string, string){
        //var inter []byte

        fromAccountInfo := db.AccountHexGet(core_arg.Db, address)

        fromBalance := new(big.Int)
        fromBalance.SetString(fromAccountInfo.Balance, 10)
        
        feeAccountInfo := db.AccountHexGet(core_arg.Db, params.Chain().Version.Eleve["dev"].FeeAddress)

        feeBalance := new(big.Int)

        //fromMinusBalance := new(big.Int)

        transFeeBalance := new(big.Int)

        feeBalance.SetString(feeAccountInfo.Balance,10)

        transFeeBalance = params.Chain().Version.Eleve["dev"].Fee

        if(fromBalance.Cmp(transFeeBalance)<0){
                return false, "balance not enough",""
        }
        inter := db.StringHexGet(core_arg.NameDb, name)
        if(inter == ""){
                db.StringHexPut(core_arg.NameDb, address, name)
                db.StringHexPut(core_arg.NameDb, name, address)
                //fmt.Println(address, name)
        fmt.Println(address,name)
        x := db.StringHexGet(core_arg.NameDb, address)
        y := db.StringHexGet(core_arg.NameDb, name)
        fmt.Println("address:",x,"name:",y)

                return true,"Create success",""
        }
        return false,"name is used",""
}
