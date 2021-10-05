package main

import (
	"fmt"
	"lurcury/account"
	"lurcury/core/block"
	"lurcury/types"
	//"math/big"
)

func InitBlock()(types.BlockJson){
        b:=block.NewBlock("sue",
        0,
        "fea4910f5d3e2d3af187cec5b8d8b1cfe99a9f5545ba50495bd42f4bae234b3a",
        0,
        0,
        "mogotisa",
        )
	return b
}

func InitAccount(core_arg types.CoreStruct, genesis types.BlockJson)(bool){
	for i:=0; i<len(genesis.Allocate); i++{
		account.GenesisAccount(
			core_arg, 
			genesis.Allocate[i].Address, 
			genesis.Allocate[i].Amount,
			genesis.Allocate[i].TokenName,
			genesis.Allocate[i].TokenBalance,
		)
		fmt.Println(
			genesis.Allocate[i].Address,
			genesis.Allocate[i].Amount,
			genesis.Allocate[i].TokenName,
			genesis.Allocate[i].TokenBalance,
		)
	}
	return true

}

