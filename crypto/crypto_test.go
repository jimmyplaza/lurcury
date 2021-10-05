package crypto

import (
	//"encoding/hex"
	eddsa "lurcury/crypto/eddsa"
	"fmt"
	"testing"
	//"reflect"
)

func TestCrypto(t *testing.T){
	x := CICKeyToAddress_hex("32df07f723d4099f2b97699a8edbd4effe0260d0bca90cec1c144c4e188f4520abefe83aba26d73ae0e31ff99938438adda65c50bfe6e483680065b2b258ada1")
	fmt.Println(x)
	y := eddsa.EddsaKeyToPublicKey([]byte("0004b80f21c57a47b074aaa34abc16f0a9c0a9a45808cdad2f267ea0bd6b8843"+"8c43c6485224b9ff3708c24bffa6f3a2bf31bda091009a19d9b964b9704899cc"))
	//var x Secp256k1
	fmt.Println(y)
	
	//x.SecpGenerateKey()
/*
	check := func(f string, got, want interface{}) {
                if !reflect.DeepEqual(got, want) {
                        t.Errorf("%s mismatch: got %v, want %v", f, got, want)
                }
        }
	address := hex.EncodeToString(KeyToAddress([]byte("9dc8a221a27d4bf0df46ba54c04e28cca51d13d10ccb1e9cb700bfa7a88a212c")))
	keccak := Keccak256([]byte("9dc8a221a27d4bf0df46ba54c04e28cca51d13d10ccb1e9cb700bfa7a88a212c"))
	x,_ := hex.DecodeString("0b1d7080dd923a7dfe42de42ee3e13feebd9c56f4c5cff6862e2d2890b4e1aba")
	fmt.Println("key:",hex.EncodeToString(KeyToAddress(x)))

	fmt.Println("hex address:",KeyToAddress_hex("2a90b5ec4c54c9da67000ff004129a5d738b86d18280270dee68833ec300ed9d"))
	check("keyToAddress()", address, "3c01f961399c0c50a7ad37778eee8a20e1c1bb32")
	check("keccak256()", keccak,"7a5189acc8bb077a3e97489661da6facdbdbf805c8b69c2ece50135c2f0dbf74")
*/
}

