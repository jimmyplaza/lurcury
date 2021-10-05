package types

type HistoryData struct {
        Nonce int `json:"nonce"`
        Balance map[string]string `json:"balance"`
        Transaction []TransactionOld `json:"transactions"`
        Address string `json:"address"`
        Urlpath string `json:"urlpath"`
}

type TransactionOld struct {
        Type string `json:"type"`
        To string `json:"to"`
        Out map[string]string `json:"out"`
        Timestamp int64 `json:"timestamp"`
        PublicKey string `json:"publickey"`
        Nonce int `json:"nonce"`
        Fee string `json:"fee"`
        From string `json:"from"`
        Sign string `json:"sign"`
        Input string `json:"input"`
        Txid string `json:"txid"`
}

type BroadcastPy struct {
        Method string `json:"method"`
        Result bool `json:"result"`
        Message string `json:"message"`
	Txid string `json:"txid"`
}

type SignTransactionPy struct {
        Method string `json:"method"`
        Message string `json:"message"`
        Param []TransactionOld `json:"param"`
	Txid string `json:"txid"`
}


