package secp256k1

import (
	//"bytes"
	"crypto/ecdsa"
	"crypto/sha256"
	//"errors"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"

	"github.com/btcsuite/btcutil/base58"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	secp "github.com/skycoin/skycoin/src/cipher/secp256k1-go"
	"golang.org/x/crypto/ripemd160"
	//"reflect"
)

func SecpSign2(pri []byte, msg []byte) []byte {
	re, err := secp256k1.Sign(msg, pri)
	if err != nil {
		fmt.Println(err)
	}
	return re
}

func SecpVerify2(pub []byte, msg []byte, sign []byte) bool {
	//x,_ :=  RecoverPubkey(msg, sign)
	//return reflect.DeepEqual(x, pub)
	return secp256k1.VerifySignature(pub, msg, sign[:64])
}

func SecpSign(pri []byte, msg []byte) []byte {
	privateKey, _ := getEcdsaKey(hex.EncodeToString(pri))
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, msg)
	if err != nil {
		fmt.Println(err)
	}
	re := []byte(r.String() + "x" + s.String())
	return re
}

func RecoverPubkey(msg []byte, sig []byte) ([]byte, error) {
	return secp256k1.RecoverPubkey(msg, sig)
}

func SigToRS(sig []byte) (*big.Int, *big.Int) {
	str := string(sig[:])
	lo := strings.Index(str, "x")
	//fmt.Println(str)
	//fmt.Println(lo)
	//fmt.Println(str[:lo])
	r, _ := new(big.Int).SetString(str[:lo], 10)
	s, _ := new(big.Int).SetString(str[lo+1:], 10)
	//fmt.Println(r)
	return r, s
}

func SecpVerify(pub []byte, msg []byte, sign []byte) bool {
	x, y := elliptic.Unmarshal(secp256k1.S256(), pub)
	var pubkey = &ecdsa.PublicKey{
		Curve: secp256k1.S256(),
		X:     x, //fromBase16(pub[:64]),
		Y:     y, //fromBase16(pub[64:]),
	}
	r, s := SigToRS(sign)
	valid := ecdsa.Verify(pubkey, msg, r, s)
	return valid
}

func fromBase16(base16 string) *big.Int {
	i, _ := new(big.Int).SetString(base16, 16)
	return i
}

func PubToEc(pub string) *ecdsa.PublicKey {
	var pubkey = &ecdsa.PublicKey{
		Curve: elliptic.P256(),
		X:     fromBase16(pub[:64]),
		Y:     fromBase16(pub[64:]),
	}
	return pubkey
}

func getEcdsaKey(randKey string) (*ecdsa.PrivateKey, error) {
	return crypto.HexToECDSA(randKey)
}

func Sha256(target []byte) []byte {
	h := sha256.New()
	h.Write(target)
	r2 := h.Sum(nil)
	return r2
}

func Ripemd160(target []byte) []byte {
	s := ripemd160.New()
	s.Write(target)
	r := s.Sum(nil)
	return r
}

func pubToAdd(btcpub []byte) string {
	head := make([]byte, 1)
	s1 := Sha256(btcpub)
	r1 := Ripemd160(s1)
	head[0] = 0x00
	r2 := append(head[:], r1...)
	s2 := Sha256(Sha256(r2))
	r3 := append(r2, s2[0:4]...)
	return base58.Encode(r3)
}

func Eth_pubToAdd(pub []byte) string {
	//pub := priv.PublicKey
	//ecdsaPubBytes := elliptic.Marshal(secp256k1.S256(), pub.X, pub.Y)
	addressBytes := crypto.Keccak256(pub[1:])[12:]
	addressTarget := hex.EncodeToString(addressBytes)
	return "0x" + addressTarget
}

func PubkeyToEC(pub []byte) (*ecdsa.PublicKey, error) {
	return crypto.UnmarshalPubkey(pub)
}

func Btc_pubToAdd(pubk []byte) (string, string) {
	//priv, _ := crypto.GenerateKey()
	pub, _ := PubkeyToEC(pubk)
	ecdsaPubBytes := elliptic.Marshal(secp256k1.S256(), pub.X, pub.Y)
	pub.Y.Rem(pub.Y, big.NewInt(2))
	var btcpubc []byte
	var btcpub []byte
	head := make([]byte, 1)
	if pub.Y.Cmp(big.NewInt(0)) == 0 {
		head[0] = 0x02
		btcpubc = append(head[:], ecdsaPubBytes[1:33]...)
	} else {
		head[0] = 0x03
		btcpubc = append(head[:], ecdsaPubBytes[1:33]...)
	}
	head[0] = 0x04
	btcpub = append(head[:], ecdsaPubBytes[1:65]...)

	return pubToAdd(btcpub), pubToAdd(btcpubc)
}

func Cic_pubToAdd(pubk []byte) string {
	return hex.EncodeToString(Sha256([]byte(hex.EncodeToString(pubk[1:]))))[24:64]
}

func SecpGenerateKey() (priv []byte, pub []byte) {
	x, y := secp.GenerateKeyPair()
	return y, x
}

/*
func SecpSign(pri []byte,msg []byte) ([]byte){
	//Sign(msg []byte, seckey []byte) ([]byte, error)
	x,y := secp256k1.Sign(msg,pri)
	if(y!=nil){
		fmt.Println(y)
	}
	return x
}

func SecpVerify(pub []byte, msg []byte, sign []byte)(bool){
	return secp256k1.VerifySignature(pub, msg, sign)
}
*/

func SecpKeyToPublicKey(key []byte) []byte {
	x, _ := getEcdsaKey(hex.EncodeToString(key))
	pub := x.PublicKey
	ecdsaPubBytes := elliptic.Marshal(secp256k1.S256(), pub.X, pub.Y)
	return ecdsaPubBytes //hex.EncodeToString(ecdsaPubBytes[1:])
}

func SecpPubToAddress(key []byte) []byte {
	var hashChannel = make(chan []byte, 1)
	sv := sha256.Sum256(key)
	hashChannel <- sv[:]
	return <-hashChannel
}

/*
func main(){
	x, y := secp.GenerateKeyPair()
	fmt.Println(x,y)
	z := SecpSign(y, []byte("123"))
	fmt.Println(z)
	d := SecpVerify(x,[]byte("123"),z)
	fmt.Println(d)
}
*/
