package transaction

import (
	"encoding/hex"
	"lurcury/types"
	"time"
)
func NewTran(crypto string)(types.TransactionJson){
        trans := types.TransactionJson{
                Balance: "1",
                To: "5834388a0f62d75bd059300fa02c1a938a301007",
                Nonce: 210,
                Fee: "10",
                Type: "a64",
                Input: "ccaacc",
		Crypto: crypto,
        }
	return trans
}

func ExpTransaction()(types.TransactionJson){
        re := NewTransaction(
                "1",
		"264411884d6d2aca8ca2d2a77c9dc95ffdcee521",
		0,
                "100",//*big.NewInt(1),
                "none",
				"none",
        )
        a,_ := hex.DecodeString("ab70ef5f36dbfd9e403ed4ffd5b1c51dc7ce761ee21c8dc72570c6d73bb9412b0b1d7080dd923a7dfe42de42ee3e13feebd9c56f4c5cff6862e2d2890b4e1aba")
        result := SignTransaction(a,re)
        //fmt.Println("re",result.Tx)
        re2 := EncodeTransaction(result)
        //fmt.Println("re2",re2)
        bb := DecodeTransaction(re2)
        //fmt.Println("bb",bb)
        return bb
}

func ExpTokenTransaction()(types.TransactionJson){
        re := NewTokenTransaction(
                "1",
		"264411884d6d2aca8ca2d2a77c9dc95ffdcee521",
		"deh",
                "",//*big.NewInt(1000),
                1,
                "",//*big.NewInt(1),
                "def",
                "none",
        )
        a,_ := hex.DecodeString("ab70ef5f36dbfd9e403ed4ffd5b1c51dc7ce761ee21c8dc72570c6d73bb9412b0b1d7080dd923a7dfe42de42ee3e13feebd9c56f4c5cff6862e2d2890b4e1aba")
        result := SignTransaction(a,re)
	//fmt.Println("re",result.Tx)
        re2 := EncodeTransaction(result)
	//fmt.Println("re2",re2)
        bb := DecodeTransaction(re2)
	//fmt.Println("bb",bb)
	return bb
}

func ExpBnnTransaction()(types.TransactionJson){
        out := []types.TransactionOut{
                {Balance: "0", Token: "bnn", Vout: 0},
        }
	protocol := ExpStation()
        trans := types.TransactionJson{
                Balance: "0",
                Out: out,
                To: "0x",
                Nonce: 0,
                Fee: "0",
                Type: "bnn",
                Input: "0fa05fzx003topyx003123",
		Protocol:protocol,
        }
	return trans 
}

func ExpLockTransaction()(types.TransactionJson){
        lock := types.TimeLock{
		Amount : "100",
		Time : time.Now(),
        }
        //protocol := ExpStation()
        trans := types.TransactionJson{
                Balance: "0",
                //Out: out,
                To: "0x",
                Nonce: 0,
                Fee: "0",
                Type: "unlock",
		TimeLock: lock,
                //Input: "",
                //Protocol:protocol,
		PublicKey:"2a90b5ec4c54c9da67000ff004129a5d738b86d18280270dee68833ec300ed9d",
        }
        return trans
}

