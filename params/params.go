package params

import (
	//"fmt"
	//"flag"
	"lurcury/types"
	"math/big"
	"os"
)

func Chain() *types.ChainConfigStructure {
	/*
	           datadir := flag.String("datadir", "../dbdata", "Data dir")
	   	port := flag.String("port", "9000", "port ")
	           chain := flag.String("chain", "BNN", "chain name ")
	           flag.Parse()
	   	fmt.Println(*chain, *datadir)
	*/
	d := &types.VersionData{
		Fee:        big.NewInt(100000000000000000),
		FeeAddress: "228533c28d5b25c7d9973afd53cb57063b13fdd0",
		// UsdtAddress: "cef49c6218ff7465e054ac25f4769c651014f482",
		UsdtAddress: os.Getenv("USDTNADDR"),
		FeeToken:    "def",
		BlockSpeed:  1,
	}
	v := &types.Version{
		Sue:   d,
		Eleve: make(map[string]*types.VersionData),
	}
	v.Eleve["dev"] = d
	v.Eleve["prod"] = d
	s := &types.ChainConfigStructure{
		Hash:    "fea4910f5d3e2d3af187cec5b8d8b1cfe99a9f5545ba50495bd42f4bae234b3a",
		Id:      101,
		V:       "0",
		Version: v,
		//Port: *port,
		//Datadir: *datadir,
		//ChainName: *chain,
	}
	return s
}
