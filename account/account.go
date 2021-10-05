package account

import (
	"encoding/hex"
	"fmt"
	crypto "lurcury/crypto"
	eddsa "lurcury/crypto/eddsa"
	"lurcury/db"
	"lurcury/types"
)

func GenesisAccount(core_arg types.CoreStruct, address string, balance string, tokenName []string, tokenBalance []string) bool {
	s := types.AccountData{
		Nonce:   0,
		Balance: balance,
		Token:   map[string]string{"kaman": "0"},
	}
	for i := 0; i < len(tokenName); i++ {
		s.Token[tokenName[i]] = tokenBalance[i]
	}
	s.Token["wen"] = "300000000000000"
	db.AccountHexPut(core_arg.Db, address, s)
	return true
}

func NewAccount() (string, string, string) {
	pri, pub := eddsa.EddsaGenerateKey()
	fmt.Println("pri: ", hex.EncodeToString(pri))
	fmt.Println("pub: ", hex.EncodeToString(pub))
	addr := crypto.KeyToAddress(pub)
	return hex.EncodeToString(pri), hex.EncodeToString(pub), (hex.EncodeToString(addr))
}

func Account_exp() types.AccountData {

	s := types.AccountData{
		Address:     "264411884d6d2aca8ca2d2a77c9dc95ffdcee529",
		Nonce:       0,
		Balance:     "1000000000000000000",
		Token:       make(map[string]string),
		Transaction: []types.TransactionJson{},
	}

	s.Token["def"] = "1000000"
	s.Token["deh"] = "1000000"

	//fmt.Println(b,s)
	return s
}

func InitAccount(address string, tokenName string, amount string) types.AccountData {
	s := types.AccountData{
		Address:     address,
		Nonce:       0,
		Balance:     "0",
		Token:       make(map[string]string),
		Transaction: []types.TransactionJson{},
	}
	s.Token[tokenName] = amount
	return s
}
