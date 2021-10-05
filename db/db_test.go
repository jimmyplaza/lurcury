package db

import (
        //"encoding/hex"
        "fmt"
	//"encoding/json"
	"lurcury/types"
	//"reflect"
        "testing"
	//"time"
)

func TestInterface(t *testing.T){
	db := OpenDB("./test")
	account_tmp := types.TransactionOut{Balance: "0", Token: "bnn", Vout: 0}
	FaceHexPut(db, account_tmp.Token, account_tmp)
	var inter types.TransactionOut
	FaceHexGet(db, account_tmp.Token,&inter)//.([]types.TransactionOut)
	//fmt.Println(add[])

        //var inter interface{}
        //json.Unmarshal(add, &inter)

	fmt.Println(inter.Token)//.([]types.TransactionOut))
	//fmt.Println("TokenName:",add[0].Token,"TokenBalance:",add[0].Balance)
}


func TestAccount(t *testing.T){
        //check := func(f string, got, want interface{}) {
        //        if !reflect.DeepEqual(got, want) {
        //                t.Errorf("%s mismatch: got %v, want %v", f, got, want)
        //        }
        //}
    //db := OpenDB("../dbdata")
	//fmt.Print(db)
	//t1 := time.Now()
/*
	for i :=1; i <=10000; i++{
		db.Put( []byte("keyddd"), []byte("11"),nil)
	}
*/
/*
	fmt.Println("put10000:",time.Now().Sub(t1))
	t2 := time.Now()
	for i2 :=1; i2 <=10000; i2++{
		f,err := db.Get([]byte("keyddd"),nil)
		if(err != nil){
			fmt.Println(f)
		}
	}
	fmt.Println("get10000:",time.Now().Sub(t2))
	f,_ := db.Get([]byte("keydd1"),nil)
	//d := string(f)
        fmt.Println(f)
	//check("get:","11",d)
*/
	/*
	account_tmp := account.Account_exp()
	AccountHexPut(db, account_tmp.Address, account_tmp)
	fmt.Println(AccountHexGet(db, "gx5ee464a101d58877f00957eff452c148e7f75833"))
	account_tmp.Address = "gx"
	AccountHexPut(db, "gx5ee464a101d58877f00957eff452c148e7f75830", account_tmp)
	fmt.Println("33",AccountHexGet(db, "gx5ee464a101d58877f00957eff452c148e7f75833"))
	fmt.Println("30",AccountHexGet(db, "1"))
	db.Put( []byte("keydd1"), []byte("11"),nil)
	b,_ := db.Get([]byte("keydd1"),nil)
	db.Put( []byte("keydd2"), []byte("12"),nil)
	fmt.Println("b:",b)
	fff,_ := db.Get([]byte("keydd1"),nil)
	fmt.Println(fff)
        keyf,_ := hex.DecodeString("aaa")
        keyd,_ := hex.DecodeString("fff")
	fmt.Println(keyf,keyd)
	*/
}
