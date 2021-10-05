package types

//"math/big"
//"time"

type TransactionOut struct {
	Token   string `json:"token"`
	Balance string/*big.Int*/ `json:"balance"`
	Vout    int `json:"vout"`
}

type SignTimeLock struct {
	Type         string `json:"type"`
	Amount       string `json:"amount"`
	Time         int64  `json:"time"`
	EndTime      int64  `json:"endTime"`
	UnlockAmount string `json:"unlockAmount"`
	IntervalTime int    `json:"intervalTime"`
}

type TransactionJson struct {
	BlockHash string           `json:"blockHash"`
	Tx        string           `json:"tx"`
	Version   string           `json:"version"`
	From      string           `json:"from"`
	To        string           `json:"to"`
	Balance   string           `json:"balance"`
	Out       []TransactionOut `json:"out"`
	//map[string]string
	Nonce     int `json:"nonce"`
	Fee       string/*big.Int*/ `json:"fee"`
	Type      string      `json:"type"`
	Input     string      `json:"input"`
	Sign      string      `json:"sign"`
	Crypto    string      `json:"crypto"`
	PublicKey string      `json:"publicKey"`
	Protocol  interface{} `json:"protocol"`
	TimeLock  TimeLock    `json:"timeLock"`
	Timestamp int64       `json:"timestamp"`
	Gas       string      `json:"gas"`
}

type SignBodys_Struct struct {
	Type       string           `json:"type"` //ex:cic
	Fee        string           `json:"fee"`
	Address    string           `json:"address"`
	Out        []TransactionOut `json:"out"` //ex:ci1,ci2
	Balance    string           `json:"balance"`
	Nonce      int              `json:"nonce"`
	Input      string           `json:"input"`
	PrivateKey string           `json:"privatekey"`
	Crypto     string           `json:"crypto"`
	Timelock   SignTimeLock     `json:"timelock"`
	Protocol   interface{}      `json:"protocol"`
}
