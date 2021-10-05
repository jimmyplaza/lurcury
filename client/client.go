package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"lurcury/core/genesis"
	"lurcury/core/transaction"
	"lurcury/db"
	lhttp "lurcury/http"
	"lurcury/params"
	"lurcury/types"
	"net/http"
	"strconv"
	"time"

	"os"
	"os/signal"
	"syscall"
)

func initchain(core_arg types.CoreStruct, chainName string) {
	tmpBlock := genesis.GenesisBlock(chainName)
	InitAccount(core_arg, tmpBlock)
	db.BlockHexPut(core_arg.Db, tmpBlock.Hash, tmpBlock)
	db.BlockNumberPut(core_arg.Db, "0", tmpBlock.Hash)
	db.BlockNumberPut(core_arg.Db, "kaman", "0")
}

func blockSync(core_arg types.CoreStruct, blocks types.BlockJson) {
	db.BlockNumberPut(core_arg.Db, "kaman", strconv.Itoa(blocks.BlockNumber))
	db.BlockHexPut(core_arg.Db, blocks.Hash, blocks)
	db.BlockNumberPut(core_arg.Db, strconv.Itoa(blocks.BlockNumber), blocks.Hash)
}

func main() {

	fmt.Println("-datadir string, -init bool , -port string, -chain string, max params three")
	datadir := flag.String("datadir", "CIC", "Data dir")
	port := flag.String("port", "9006", "port ")
	chain := flag.String("chain", "CIC", "chain name ")
	peers := flag.String("peer", "http://192.168.51.203:9006", "peer ip ")
	model := flag.String("model", "1", "model")
	flag.Parse()

	config := params.Chain()
	config.Port = *port
	config.Datadir = *datadir
	config.ChainName = *chain
	config.Peers = append(config.Peers, *peers)

	core_arg := &types.CoreStruct{}
	core_arg.Config = config
	core_arg.Db = db.OpenDB("../../cclient" + config.Datadir)
	core_arg.NameDb = db.OpenDB("../../cclient" + config.Datadir + "/Name")
	core_arg.Model = *model
	num := db.BlockNumberGet(core_arg.Db, "kaman")
	if num == "" {
		fmt.Println("num is null, do initchain()")
		initchain(*core_arg, *datadir) // "CIC")//*datadir)
	} else {
		fmt.Println("num is not null, print num(block num)")
		//num := db.BlockNumberGet(core_arg.Db, "kaman")
		//hex := db.BlockNumberGet(core_arg.Db, num)
		//tmpBlock = db.BlockHexGet(core_arg.Db, hex)
		log.Println(num)
	}

	go lhttp.Server(core_arg)

	// Broadcast to Peer node  xxx/broadcast
	go Cast(core_arg)

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	check := 0
	go func() {
		<-c
		fmt.Println("cleanup")
		//time.Sleep(1 * time.Second)
		check = 1
	}()

	for {
		num = db.BlockNumberGet(core_arg.Db, "kaman")
		targettmp, _ := strconv.Atoi(num)
		target := strconv.Itoa(targettmp + 1)
		log.Println("sync block:", target)
		resp, err := http.Get(config.Peers[0] + "/getBlockbyID?blockID=" + target)
		resp2, err2 := http.Get(config.Peers[0] + "/getBlockNum")
		if err != nil {
			fmt.Println(err)
			fmt.Println(err2)
			time.Sleep(1 * time.Second)
		} else {
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println(err)
				time.Sleep(1 * time.Second)
			}
			body2, err2 := ioutil.ReadAll(resp2.Body)
			fmt.Println("getBlockNum: ", string(body2))
			if err2 != nil {
				fmt.Println(err2)
				time.Sleep(1 * time.Second)
			}
			if string(body) != "no block hash" {
				var reqBlock types.BlockJson
				json.Unmarshal(body, &reqBlock)
				fmt.Println("===============================")
				fmt.Printf("%#v", reqBlock)
				blockSync(*core_arg, reqBlock)
				for i := 0; i < len(reqBlock.Transaction); i++ {
					// _, err := transaction.TransactionProtocol(*core_arg, reqBlock.Transaction[i])
					_, err := transaction.TransactionProtocol(core_arg, reqBlock.Transaction[i])
					log.Println(reqBlock.Transaction[i].Tx, "sync state:", err)
					if err != "success" {
						//os.Remove(*datadir)
						//initchain(*core_arg,*datadir)
						os.Exit(0)
					}
				}
			} else {
				mainblocknum, _ := strconv.Atoi(string(body2))
				if targettmp < mainblocknum {
					db.BlockNumberPut(core_arg.Db, "kaman", target)
				} else {
					time.Sleep(10000 * time.Millisecond)
				}
			}
		}
		if check == 1 {
			os.Exit(1)
		}
	}
}
