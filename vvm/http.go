package main

import (
        "github.com/gin-gonic/gin"
	"lurcury/types"
	"github.com/gin-contrib/cors"
)
var core *types.CoreStruct

func HttpRun(core_arg *types.CoreStruct) {
	core = core_arg
        r := gin.Default()
        
	config := cors.DefaultConfig()
        config.AllowAllOrigins = true
        r.Use(cors.New(config))
	r.Use(cors.Default())

        r.GET("/ping", ping)
	r.GET("/txsearch", txSearch)
	r.GET("/addresssearch", addressSearch)
        r.POST("/esGas", esGas)
        r.Run(":5214")
}

func ping(c *gin.Context){
        c.JSON(200, gin.H{
                "message": "pong",
        })
}

//curl -X POST -H 'Content-Type: application/json' -d '{"from": "d40ae259ba4c696c441e2a7f9e3fd175f1899bbc","to": "0xfde46C0Fb4172274c533fBC4de354488f1587C5b","balance": "0","nonce": 9,"input":"0xb518a776", "type":"VvmDCall"}' http://127.0.0.1:5214/esGas
func esGas(c *gin.Context){
	var data types.TransactionJson
        c.BindJSON(&data)
	a,b,d,_ := Drill_Run_decode(core.Evm, data)
        c.JSON(200, gin.H{
                "ret": a,
		"evmContractAddress": b,
		"usingGas":d,
        })
}

func txSearch(c *gin.Context){
	tx := c.Query("tx")
	tran := core.ContractTx[tx]
        c.JSON(200, tran)
}

func addressSearch(c *gin.Context){
        address := c.Query("address")
        tran := core.ContractTx[address]
        c.JSON(200, tran)
}


