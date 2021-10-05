package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/syndtr/goleveldb/leveldb"
	//"lurcury/core/block"
	//"lurcury/core/transaction"
	//"lurcury/params"
	"time"
)

type ChainConfigStructure struct {
	Id        int64
	Hash      string
	V         string
	Version   *Version
	Port      string
	Datadir   string
	ChainName string
	Peers     []string
}

type VersionData struct {
	Fee              *big.Int
	FeeAddress       string
	UsdtAddress      string
	FeeToken         string
	BlockSpeed       int
	BlockTransaction int
	Consensus        string
}

type Version struct {
	Sue   *VersionData
	Eleve map[string]*VersionData
}

type TimeLock struct {
	Type                          string    `json:"type"`
	Amount                        string    `json:"amount"`
	Time                          time.Time `json:"time"`
	EndTime                       time.Time `json:"endTime"`
	UnlockAmount                  string    `json:"unlockAmount"`
	IntervalTime/*time.Time*/ int           `json:"intervalTime"`
}

type BalanceData struct {
	Token   string
	Balance string //*big.Int
}

type AccountData struct {
	Address string
	Nonce   int
	Balance string
	//Token []BalanceData
	Token       map[string]string
	Transaction []TransactionJson
	TimeLock    []TimeLock
}

type CoreStruct struct {
	Test                      string
	PendingBlockTransaction   []TransactionJson
	PendingBalanceTransaction []TransactionJson
	PendingNonceTransaction   []TransactionJson
	PendingSignTransaction    []TransactionJson
	PendingTransaction        []TransactionJson
	Db                        *leveldb.DB
	NameDb                    *leveldb.DB
	Config                    *ChainConfigStructure
	PendingBlock              []BlockJson
	Model                     string
	Evm                       *vm.EVM
	ContractTx                map[string]ContractInfo
	ContractAddress           map[string][]ContractInfo
}

type ContractInfo struct {
	Ret     string
	Address string
	Err     error
	Gas     uint64
}
