package secp256k1
import (
	"encoding/hex"
	//"crypto/elliptic"
	//"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"fmt"
	"testing"
	//"github.com/ethereum/go-ethereum/crypto"
	//"crypto/rand"
	//"crypto/ecdsa"
)


func TestCrypto(t *testing.T){

        p:="98123fa5d05103e0e1ad7951aa280c0e98aee95ed4c2594333b477045e6cbc89"

	d,_ := hex.DecodeString(p)
/*
	r := SecpSign(d, []byte("123111111111111111111111111111111111111111111111"))
	pub := "042ccf85a1706ae8583e2839177db014d97df1e7972712c6d5cc8130f3fca652d5bdeaa3cabe80621a9747e9cc2ffde40b3df45e63ac22a9a2e71b17b070736de5"
	dd,_ := hex.DecodeString(pub)
	fmt.Println(Cic_pubToAdd(dd))
	re := SecpVerify(dd,[]byte("123111111111111111111111111111111111111111111111"),r)
	fmt.Println(re)
*/

	msg,_ := hex.DecodeString("1111111111111111111111111111111111111111111111111111111111111111")
	r2 := SecpSign2(d, msg)
	fmt.Println("sign:", hex.EncodeToString(r2))
	x,_ := RecoverPubkey(msg,r2)
	//fmt.Println(r2[:64])
	re2 := SecpVerify2(x, msg,r2)
	fmt.Println("result:",re2)
	fmt.Println(hex.EncodeToString(x))

	//re3 := SecpVerify2(x, msg, r2)
	//fmt.Println(re3)

	//privateKey,_ := getEcdsaKey(p)
	//fmt.Println(privateKey.PublicKey.X)
	//fmt.Println(privateKey.PublicKey)

/*
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, []byte("11111111111111111111111111111111"))
	if err != nil {
		panic(err)
	}
	
	valid := ecdsa.Verify(&privateKey.PublicKey, []byte("11111111111111111111111111111111"), r, s)
	fmt.Println(valid)
*/
/*

	//fmt.Println(accountGenerate())
	p:="98123fa5d05103e0e1ad7951aa280c0e98aee95ed4c2594333b477045e6cbc89"
	x,e := getEcdsaKey(p)
	fmt.Println(e)
	fmt.Println("key:",x)
	//fmt.Println("y:",y)
        pub := x.PublicKey
	//fmt.Println(pub.X)
	//fmt.Println(pub.Y)
	ecdsaPubBytes := elliptic.Marshal(secp256k1.S256(), pub.X, pub.Y)
	fmt.Println("pub:",hex.EncodeToString(ecdsaPubBytes))
	//fmt.Println(len(ecdsaPubBytes))
	fmt.Println("cic address",hex.EncodeToString(Sha256([]byte(hex.EncodeToString(ecdsaPubBytes[1:]))))[24:64])
	etha := Eth_pubToAdd(ecdsaPubBytes)
	fmt.Println(etha)
	btca, btcac := Btc_pubToAdd(ecdsaPubBytes)
	fmt.Println(btca)
	fmt.Println(btcac)
	eee,_ := hex.DecodeString(p)
	pr, pb := crypto.GenerateKey()
	ddd := SecpSign(pr,[]byte("11111111111111111111111111111111"))

	rr := SecpVerify(pb,[]byte("11111111111111111111111111111111"),ddd)
	fmt.Println(rr)
*/
	/*
        x, y := SecpGenerateKey()
        fmt.Println(x,y)
        z := SecpSign(x, []byte("123"))
        fmt.Println(z)
        d := SecpVerify(y,[]byte("123"),z)
        fmt.Println(d)
	k := SecpKeyToPublicKey(x)
	e := SecpPubToAddress(k)
	fmt.Println("priv:",hex.EncodeToString(x))
	fmt.Println("pub:",hex.EncodeToString(y))
	fmt.Println("pub:",hex.EncodeToString(k))
	fmt.Println("add:",hex.EncodeToString(e)[26:64])
	fmt.Println(SecpVerify(k,[]byte("123"),z))
	*/
}
