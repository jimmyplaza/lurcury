package main

import (
	"crypto/ecdsa"
	"fmt"
	"crypto/elliptic"
	"math/big"
)

func fromBase16(base16 string) *big.Int {
	i, _ := new(big.Int).SetString(base16, 16)
/*
	if !ok {
		log.Fatalln("trying to convert from base16 a bad number: ", base16)
	}
*/
	return i
}

func pub(pub string)(*ecdsa.PublicKey){
        var pubkey = &ecdsa.PublicKey{
                Curve: elliptic.P256(),
		X:     fromBase16(pub[:64]),
		Y:     fromBase16(pub[64:]),
        }
	return pubkey
}

func main(){
	var pubkey = &ecdsa.PublicKey{
		Curve: elliptic.P256(),
		X:     fromBase16("3bac7e95a003264cc075a2ba8d4e949862acd755d49094ad8d28bd0d56299dc6"),
		Y:     fromBase16("5c6a5b3810181d82f5eb1be32c9cd8d6c387fcb06fed530d749e3997eb22bd8c"),
	}
	fmt.Println("2ccf85a1706ae8583e2839177db014d97df1e7972712c6d5cc8130f3fca652d5bdeaa3cabe80621a9747e9cc2ffde40b3df45e63ac22a9a2e71b17b070736de5"[:64])
	fmt.Println("2ccf85a1706ae8583e2839177db014d97df1e7972712c6d5cc8130f3fca652d5bdeaa3cabe80621a9747e9cc2ffde40b3df45e63ac22a9a2e71b17b070736de5"[64:])
	fmt.Println(pubkey)
}
