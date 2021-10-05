package main

import (
	//"context"
	"encoding/hex"
	"encoding/json"
	"log"
	"github.com/ethereum/go-ethereum/common"
	//ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/ethereum/go-ethereum/core/vm"
	"lurcury/types"
	"math/big"
	"strconv"
	"time"
)

type Params struct {
        context            vm.Context
        nonce              uint64
        executorRawAddress *common.Address
        amount             *big.Int
        contract           *common.Address
        gas                uint64
        data               []byte
        from               common.Address
        gaslimit           uint64
        gasprice           uint64
        blocknumber        uint64
	config             vm.Config
	chainConfig        *params.ChainConfig
	input              string
}

func NewContext() Params {
	var param Params
	canTransferFunc := func(db vm.StateDB, addr common.Address, amount *big.Int) bool {
		return true
	}
	transferFunc := func(db vm.StateDB, sender, recipient common.Address, amount *big.Int) {
	}
	getHashFunc := func(height uint64) common.Hash {
		return common.HexToHash("0x1A0b54D5dc17e0AadC383d2db43B0a0D3E029c4c")//common.Hash{}
	}
        fromAddress := common.BytesToAddress([]byte("0x"))
	coinbase := common.BytesToAddress([]byte("0x5A0b54D5dc17e0AadC383d2db43B0a0D3E029c4c"))
	param.context = vm.Context{
		CanTransfer: canTransferFunc,
		Transfer:    transferFunc,
		GetHash:     getHashFunc,
		Origin:      fromAddress,
		Coinbase:    coinbase,
		BlockNumber: new(big.Int).SetUint64(10000000),
		Time:        new(big.Int).SetInt64(1),
		GasLimit:    uint64(50000000),
		GasPrice:    new(big.Int).SetInt64(100),
	}
	return param
}

func Create(evm *vm.EVM,params Params, code string)([]byte, common.Address, uint64, error){
	data, err := hexutil.Decode(code)
	if(err != nil){
		return []byte(""), common.BytesToAddress([]byte("0x")), 0, err
	}
	con := NewContext()
	executor := vm.AccountRef(con.context.Origin)
	ret, evmContractAddress, remainingGas, err := evm.Create(executor, data, params.gaslimit,new(big.Int).SetInt64(0))
	return ret, evmContractAddress, remainingGas, err
}

func Call(evm *vm.EVM, params Params, target common.Address)([]byte, uint64, error){
	input, err := hexutil.Decode(params.input)
	if(err == nil){
	        executor := vm.AccountRef(params.from)
		ret, remainingGas, err := evm.Call(executor, target,input, params.gaslimit, new(big.Int).SetInt64(0))
		return ret, remainingGas, err
	}else{
		return []byte(""), 0, err
	}
}

func Call_decode(evm *vm.EVM, params Params, target common.Address)(string,uint64,error){
	a,b,c := Call(evm, params, target)
	return hex.EncodeToString(a), b, c
}

func EstimateGas(evm *vm.EVM, params Params, target common.Address)([]byte, uint64, error){
        input, err := hexutil.Decode(params.input)
        if(err == nil){
                executor := vm.AccountRef(params.from)
                ret, remainingGas, err := evm.DrillCall(executor, target,input, params.gaslimit, new(big.Int).SetInt64(0))
                return ret, 8000000-remainingGas, err
        }else{
                return []byte(""), 0, err
        }
}

func EstimateGas_decode(evm *vm.EVM, params Params, target common.Address)(string,uint64,error){
        a,b,c := EstimateGas(evm, params, target)
        return hex.EncodeToString(a), b, c
}

func Drill_Create(evm *vm.EVM,params Params, code string)([]byte, common.Address, uint64, error){
        data, err := hexutil.Decode(code)
        if(err != nil){
                return []byte(""), common.BytesToAddress([]byte("0x")), 0, err
        }
        con := NewContext()
        executor := vm.AccountRef(con.context.Origin)
        ret, evmContractAddress, remainingGas, err := evm.DrillCreate(executor, data, params.gaslimit,new(big.Int).SetInt64(0))
        return ret, evmContractAddress, remainingGas, err
}

func Drill_Create_decode(evm *vm.EVM,params Params, code string)(string,uint64,error){
        a,_,b,c := Drill_Create(evm, params, code)
        return hex.EncodeToString(a), b, c
}

func Run(evm *vm.EVM,transaction types.TransactionJson)([]byte, common.Address, uint64, error){
	amounts,_ := new(big.Int).SetString(transaction.Balance,10)
	var fromAddress common.Address
	fromAddress = common.HexToAddress(transaction.From)
	if(transaction.Type == "VvmCreate"||transaction.Type == "a65"){
		setParams := Params{
                                nonce: uint64(transaction.Nonce),
                                from: fromAddress,
                                amount: amounts,
                                gaslimit: 8000000,
                                gasprice: 1,
                                input:transaction.Input,
		}
        	ret, evmContractAddress, remainingGas, err := Create(evm, setParams, setParams.input)
		return ret, evmContractAddress, remainingGas, err
	}else if(transaction.Type == "VvmCall"||transaction.Type == "a66"){
                setParams := Params{
                                nonce: uint64(transaction.Nonce),
                                from: fromAddress,
                                amount: amounts,
                                gaslimit: 8000000,//3000000,
                                gasprice: 1100,
                                input:transaction.Input,
                }
                ret, remainingGas, err := Call(evm,setParams, common.HexToAddress(transaction.To))
		return ret, common.HexToAddress(transaction.To), remainingGas, err
        }else{
		return []byte(""), common.HexToAddress("0x"), 0, nil
	}
}

func Drill_Run(evm *vm.EVM,transaction types.TransactionJson)([]byte, common.Address, uint64, error){
	//log.Println(transaction)
        amounts,_ := new(big.Int).SetString(transaction.Balance,10)
        var fromAddress common.Address
        fromAddress = common.HexToAddress(transaction.From)
        if(transaction.Type == "VvmDCall"||transaction.Type == "a66"){
                setParams := Params{
                                nonce: uint64(transaction.Nonce),
                                from: fromAddress,
                                amount: amounts,
                                gaslimit: 8000000,
                                gasprice: 1,
                                input:transaction.Input,
                }
                ret, remainingGas, err := EstimateGas(evm,setParams, common.HexToAddress(transaction.To))
                return ret, common.HexToAddress(transaction.To), remainingGas, err
        }else if(transaction.Type == "VvmDCreate"||transaction.Type == "a65"){
                setParams := Params{
                                nonce: uint64(transaction.Nonce),
                                from: fromAddress,
                                amount: amounts,
                                gaslimit: 8000000,
                                gasprice: 1,
                                input:transaction.Input,
                }
                ret, evmContractAddress, remainingGas, err := Drill_Create(evm, setParams, transaction.Input)
                log.Println("create result:",ret, remainingGas, err)
                return ret, evmContractAddress, remainingGas, err
        }else{
                return []byte(""), common.HexToAddress("0x"), 0, nil
        }
}

func Run_decode(evm *vm.EVM,transaction types.TransactionJson)(string, string, uint64, error){
	ret, address,gas,err := Run(evm,transaction)
	return hex.EncodeToString(ret), address.Hex(), gas, err
}

func Drill_Run_decode(evm *vm.EVM,transaction types.TransactionJson)(string, string, uint64, error){
	if(transaction.Type == "VvmDCreate"||transaction.Type == "VvmDCall"||transaction.Type == "a65"||transaction.Type == "a66"){
        	ret, address,gas,err := Drill_Run(evm,transaction)
        	return hex.EncodeToString(ret), address.Hex(), gas, err
	}else{
		return "","",0,nil
	}
}

func VvmStart(core_arg *types.CoreStruct){
	target := 0
        for{
		target = target+1
		targettmp:= strconv.Itoa(target)
		g, e := Get(core_arg.Config.Peers[0]+"/getBlockbyID?blockID="+targettmp)
		n, _ := Get(core_arg.Config.Peers[0]+"/getBlockNum")
		number, _ := strconv.Atoi(string(n))
		if(e==nil&&number>=target&&string(g)!="no block hash"){
		        var reqBlock types.BlockJson
		        json.Unmarshal(g, &reqBlock)
			log.Println("sync blockData:",reqBlock.BlockNumber)
			for _,trans := range reqBlock.Transaction{
				trans.Input = "0x"+trans.Input
				trans.To = "0x"+trans.To
				trans.From = "0x"+trans.From
				ret,addr,gas,err := Run_decode(core_arg.Evm, trans)
				if(string(ret)!=""&&err==nil){
					log.Println("ret:",ret)
					log.Println("contract address:", addr)
					log.Println("Cost:",gas)
					if(core_arg.Model == "3"){
						info := types.ContractInfo{
                                                        Ret:ret,
                                                        Address:addr,
                                                        Gas:3000000-gas,
                                                        Err:err,
                                                }
                				core_arg.ContractTx[trans.Tx] = info
						core_arg.ContractAddress[addr] = append(core_arg.ContractAddress[addr], info)
					}
				}
			}
			//log.Println(core_arg.ContractAddress)
			log.Println("sync blockData complete:",reqBlock.BlockNumber)
		}else{
			target = target-1
			time.Sleep(10000 * time.Millisecond)
		}
	}
}


