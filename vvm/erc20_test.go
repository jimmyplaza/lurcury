package main

import (
        "github.com/ethereum/go-ethereum/common"
        "github.com/ethereum/go-ethereum/core/rawdb"
        "github.com/ethereum/go-ethereum/core/state"
        "github.com/ethereum/go-ethereum/core/vm"
        "lurcury/types"
        "log"
        "testing"
)

func TestERC20(t *testing.T){
        con := NewContext()
        con.chainConfig = GetChainConfig()
        statedb,_ := state.New(common.Hash{}, state.NewDatabase(rawdb.NewMemoryDatabase()))
        evm := vm.NewEVM(con.context, statedb, con.chainConfig, con.config)
        code:="0x6080604052348015600f57600080fd5b5060988061001e6000396000f300608060405260043610603e5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663f3ccaac081146043575b600080fd5b348015604e57600080fd5b5060556067565b60408051918252519081900360200190f35b600a905600a165627a7a72305820ecd477a3fff5fd0b6b214dc1ad5951b840e808327607cc9459245ffdef7100890029"
        code2:="0x608060405260008054600160a060020a03191673c17ac11630b7aa50cac604c74a0e239a4cf236c517905534801561003657600080fd5b50610150806100466000396000f3006080604052600436106100405763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663f8a8fd6d8114610045575b600080fd5b34801561005157600080fd5b5061005a61006c565b60408051918252519081900360200190f35b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f3ccaac06040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b1580156100f357600080fd5b505af1158015610107573d6000803e3d6000fd5b505050506040513d602081101561011d57600080fd5b50519050905600a165627a7a7230582020571efea0ac0dd21e3e2aa75ab9b33dc4086261f129b009296922e8c72a9be20029"

        dtransaction := types.TransactionJson{
                Type:"VvmCreate",
                Nonce:1,
                From:"0xea674fdde714fd979de3edf0f56aa9716b898e11",
                Balance:"0",
                Input:code,
                To:"0xea674fdde714fd979de3edf0f56aa9716b898e11",
        }
        res,address,usegas,err := Run_decode(evm, dtransaction)
        log.Println("test ERC20", res, address, usegas, err)
        dtransaction2 := types.TransactionJson{
                Type:"VvmCreate",
                Nonce:1,
                From:"0xea674fdde714fd979de3edf0f56aa9716b898e11",
                Balance:"0",
                Input:code2,
                To:"0xea674fdde714fd979de3edf0f56aa9716b898e11",
        }
        res2,address2,usegas2,err2 := Run_decode(evm, dtransaction2)
        log.Println("test ERC20", res2, address2, usegas2, err2)
        etransaction := types.TransactionJson{
                Type:"VvmCall",
                Nonce:1,
                From:"0xea674fdde714fd979de3edf0f56aa9716b898e11",
                Balance:"0",
                Input:"0xf8a8fd6d",
                To:address2,
        }
        res, address, usegas,err = Run_decode(evm, etransaction)
        log.Println("ERC20 test excute", res, address, usegas, err)
}

