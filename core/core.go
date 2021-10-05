package main

import (
	"flag"
	"fmt"
	"lurcury/core/block"
	"lurcury/core/genesis"
	"lurcury/db"
	"lurcury/http"
	"lurcury/params"
	"lurcury/types"
	"time"
	/*
	   "os"
	   "os/signal"
	   "syscall"
	*/)

func main() {
	fmt.Println("-datadir string, -init bool , -port string, -chain string, max params three")
	datadir := flag.String("datadir", "../dbdata", "Data dir")
	port := flag.String("port", "9002", "port ")
	chain := flag.String("chain", "TTN", "chain name ")
	init := flag.Bool("init", false, "init ")
	flag.Parse()

	fmt.Println(*init)
	config := params.Chain()
	config.Port = *port
	config.Datadir = *datadir
	config.ChainName = *chain
	//fmt.Println("run",config)
	core_arg := &types.CoreStruct{}
	core_arg.Model = "0"
	core_arg.Config = config
	dbPath := "../../ccore" + config.Datadir
	// core_arg.Db = db.OpenDB("../../ccore" + config.Datadir) //*datadir)  // ccoreCIC
	core_arg.Db = db.OpenDB(dbPath)
	fmt.Println("===================Server leveldb: ", dbPath)
	core_arg.NameDb = db.OpenDB("../../ccore" + config.Datadir + "/Name")
	fmt.Sprintf("===============config: %V", config)
	go http.Server(core_arg)
	var tmpBlock types.BlockJson

	num := db.BlockNumberGet(core_arg.Db, "kaman")

	if num == "" {
		genesis.InitAccount(*core_arg, genesis.GenesisBlock(config.Datadir))
		tmpBlock = genesis.GenesisBlock(config.Datadir)
		db.BlockHexPut(core_arg.Db, tmpBlock.Hash, tmpBlock)
		db.BlockNumberPut(core_arg.Db, "0", tmpBlock.Hash)
		db.BlockNumberPut(core_arg.Db, "kaman", "0")
	} else {
		num := db.BlockNumberGet(core_arg.Db, "kaman")
		hex := db.BlockNumberGet(core_arg.Db, num)
		tmpBlock = db.BlockHexGet(core_arg.Db, hex)
	}
	pri := "219a634773d787cfbaf1e5c915d56b14be2a3695ed8e46bbeb01573bf211d0ef8773580834eb42a2f2ee856b029a88dfee639e27f08b1e0235f8eb04eecf4089"
	for i := 0; i != -1; {
		if len(core_arg.PendingTransaction) > 0 {
			fmt.Println("(Server) $$$$$$$$$ PendingTransaction > 0  $$$$$$$$$$$$$$$$$$$$$$")
			time.Sleep(1 * time.Second)
			ff := time.Now()
			tmpBlock = block.CreateBlockPOA(core_arg, tmpBlock, pri)
			fmt.Println(time.Now().Sub(ff))
		} else {
			fmt.Println("no transaction")
			time.Sleep(1 * time.Second)
		}
	}
}
