package crypto

import (
	//"bytes"
	"crypto/sha256"
	"encoding/hex"
	secp "lurcury/crypto/secp256k1"
	"github.com/ethereum/go-ethereum/crypto"
)

func CICKeyToAddress(key []byte)([]byte){
	return secp.Sha256([]byte(hex.EncodeToString(key[1:])))[24:64]
}

func CICKeyToAddress_hex(key string)(string){
	c,_ := hex.DecodeString(key[2:])
        return hex.EncodeToString(secp.Sha256([]byte(hex.EncodeToString(c))))[24:64]
}

func KeyToAddress(key []byte)([]byte){
	return Keccak256(key)[12:]
}

func KeyToAddress_hex(key string)(string){
	c,_ := hex.DecodeString(key)
	re := Keccak256(c)[12:]
	return hex.EncodeToString(re)//Keccak256(key)[12:]
}

func Keccak256(msg []byte)([]byte){
	return crypto.Keccak256(msg)
}

func Sha256(msg []byte)([]byte){
	h := sha256.New()
	h.Write([]byte(msg))
	bs := h.Sum(nil)
	return bs
}

func PubkeyToAddress(crypto string,publicKey string)(string){
	var address string
        switch crypto {
                case "cic":
                        address = CICKeyToAddress_hex(publicKey)
                case "secp256k1":
                        address = CICKeyToAddress_hex(publicKey)
                case "eddsa":
                        address = KeyToAddress_hex(publicKey)
                default:
	address = KeyToAddress_hex(publicKey)
	}
	return address
}
