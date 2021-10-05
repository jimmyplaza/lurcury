package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"lurcury/core/transaction"
	"lurcury/types"
	"net/http"
	"time"
)

/*
func main(){
        core_arg := &types.CoreStruct{}
        cast(core_arg)
}
*/
func Cast(core_arg *types.CoreStruct) {
	for {
		log.Println("check broadcast")
		if len(core_arg.PendingTransaction) > 0 {
			fmt.Println("$$$$$$$$$$$$ PendingTransaction > 0 $$$$$$$$$$$$")
			req := core_arg.PendingTransaction[0]
			postGo(core_arg.Config.Peers[0]+"/broadcast",
				"application/x-www-form-urlencoded", req)
			fmt.Println("Post to Peer node !! xxxx/broadcast")
			transaction.OrderDeletPendingTransaction(core_arg, 0)
		} else {
			time.Sleep(2 * time.Second)
		}
	}
}

func postGo(urls string, heads string, bodys interface{}) {
	req := bodys
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(req)
	resp, err := http.Post(urls,
		heads,
		b,
	)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	log.Println(string(body))
}
