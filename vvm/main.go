package main

import(
	"flag"
	"github.com/ethereum/go-ethereum/common"
        "github.com/ethereum/go-ethereum/core/rawdb"
        "github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/core/vm"
	"log"
	"lurcury/db"
	param "lurcury/params"
	"lurcury/types"
	"math/big"
)
func GetChainConfig() *params.ChainConfig {
	/*
        var chainConfig params.ChainConfig
        chainConfig.ConstantinopleBlock = new(big.Int).SetUint64(15000000)
        return &chainConfig
	*/

return &params.ChainConfig{
	PetersburgBlock: big.NewInt(1),
}

}


func InitConfig(core_arg *types.CoreStruct){
        log.Println("-datadir string, -init bool , -port string, -chain string, max params three")
        datadir := flag.String("datadir", "CIC", "Data dir")
        port := flag.String("port", "9006", "port ")
        chain := flag.String("chain", "CIC", "chain name ")
        peers := flag.String("peer", "http://192.168.51.212:9999", "peer ip ")
        model := flag.String("model", "3", "model")
        flag.Parse()

        config := param.Chain()
        config.Port =  *port
        config.Datadir = *datadir
        config.ChainName = *chain
        config.Peers = append(config.Peers ,*peers)

        core_arg.Config = config
        core_arg.Db = db.OpenDB("../../cclient"+config.Datadir)
        core_arg.NameDb = db.OpenDB("../../cclient"+config.Datadir+"/Name")
        core_arg.Model = *model

	con := NewContext()
	con.chainConfig = GetChainConfig()
	statedb,_ := state.New(common.Hash{}, state.NewDatabase(rawdb.NewMemoryDatabase()))
	core_arg.Evm = vm.NewEVM(con.context, statedb, con.chainConfig, con.config)
}


func main(){
        core_arg := &types.CoreStruct{
		ContractTx:make(map[string]types.ContractInfo),
		ContractAddress:make(map[string][]types.ContractInfo),
	}
        InitConfig(core_arg)
        go VvmStart(core_arg)
        HttpRun(core_arg)
}
