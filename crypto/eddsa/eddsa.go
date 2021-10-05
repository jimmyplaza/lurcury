//package main
package eddsa

import (
	//"crypto/rand"
	//"encoding/hex"
	"fmt"
	eddsa "golang.org/x/crypto/ed25519"
	"crypto/rand"
	//"crypto"
	//"time"
)

func EddsaGenerateKey()([]byte,[]byte){
	pub,priv,err := eddsa.GenerateKey(rand.Reader)
	if err != nil {
                fmt.Println("generateKey error: %s", err)
        }
	return priv,pub
}

func EddsaSign(pri []byte,msg []byte) ([]byte){
	return eddsa.Sign(pri,msg)
}

func EddsaVerify(pub []byte, msg []byte, sign []byte)(bool){
	return eddsa.Verify(pub,[]byte(msg),sign)
}

func EddsaKeyToPublicKey(key []byte)([]byte){
	return key[32:]//.Public()
}


/*
func main(){
	pub,pri,_ := eddsa.GenerateKey(rand.Reader)
	fmt.Println("publicKey:",hex.EncodeToString(pub))
	fmt.Println("privateKey:",hex.EncodeToString(pri))
	sign := eddsaSign(pri,[]byte("111111111111111111111111111111111111111111111111111111111111"))
	fmt.Println("sign:",hex.EncodeToString(sign))
	eddsaVerify(pub,[]byte("111111111111111111111111111111111111111111111111111111111111"),sign)
	fmt.Println(eddsa.PrivateKey([]byte("111111111111111111111111111111111111111111111111111111111111")).Public())
	fmt.Println(pri.Public())
}
*/
