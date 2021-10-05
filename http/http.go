package http

import (
	"fmt"
	"lurcury/http/route"
	"lurcury/types"
	"net/http"
	"time"
)

func httpSet(coreStruct *types.CoreStruct) {
	route.Router(coreStruct)
	fmt.Println(coreStruct.Config)
	err2 := http.ListenAndServe(":"+coreStruct.Config.Port /*":9000"*/, nil)
	if err2 != nil {
		fmt.Println("error:", err2)
	}

}

func httpSet2(coreStruct types.CoreStruct) {
	fmt.Println("connect port" + ":14456")
	route.Test(coreStruct)
	err2 := http.ListenAndServe(":14456", nil)
	if err2 != nil {
		fmt.Println("error:", err2)
	}

}

func Server(coreStruct *types.CoreStruct) {
	httpSet(coreStruct)
	time.Sleep(100 * time.Second)
}
