package types

import (
//        "lurcury/core/transaction"
//	"math/big"
)


type VerifierJson struct{
        Verifier string
        Sign string
        N int
}

type BlockJson struct{
        Version string `json:"version"`
        BlockNumber int `json:"blockNumber"`
        ParentHash string `json:"parentHash"`
        Nonce int `json:"nonce"`
        Transaction []TransactionJson `json:"transaction"`
        Timestamp int64 `json:"timestamp"`
        ExtraData string `json:"extraData"`
        Hash string `json:"hash"`
	Txs int `json:"txs"`
        Verifier []VerifierJson `json:"verifier"`
        Allocate []AllocateStruct `json:"allocate"`
}

type AllocateStruct struct {
        Address string
        Amount string//big.Int
	PrivateKey string
	TokenName []string
	TokenBalance []string
}


