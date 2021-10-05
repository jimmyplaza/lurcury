package transaction

import (
	"fmt"
	"lurcury/crypto"
	"lurcury/db"
	"lurcury/types"
	"strconv"
	"strings"
)

func ExpStation() types.NewsStation {
	node := []types.Content{
		{Title: "top", Intro: "go", Time: "1199", Index: "3"},
	}
	args := []string{"a", "b"}
	out := types.NewsStation{
		Name:    "0",
		Intro:   "bnn",
		Article: node,
		Picture: node,
		Node:    args,
	}
	fmt.Println(out)
	return out
}

func BnnTransactionToStation_new(trans types.TransactionJson) types.NewsStation {
	//trans.Protocol.(types.NewsStation)
	return trans.Protocol.(types.NewsStation)
}

func BnnTransactionToStation(trans types.TransactionJson) types.NewsStation {
	//取位數
	symbollength := 2
	degit := 3
	//Name
	transaction := trans.Input
	z := strings.Index(transaction, "zx")
	zl, _ := strconv.Atoi(trans.Input[z+symbollength : z+symbollength+degit])
	zv := trans.Input[z+symbollength+degit : z+symbollength+degit+zl]
	//Intro
	y := strings.Index(transaction, "yx")
	yl, _ := strconv.Atoi(trans.Input[y+symbollength : y+symbollength+degit])
	yv := trans.Input[y+symbollength+degit : y+symbollength+degit+yl]
	nodeArray := NodeinfoDecode(trans.Input, 3)

	return types.NewsStation{
		Name:  zv, //trans.Input[6:18],
		Intro: yv, //trans.Input[18:36],
		Node:  nodeArray,
		Owner: crypto.KeyToAddress_hex(trans.PublicKey),
	}
}

func NodeinfoDecode(info string, degit int) []string {
	symbollength := 2
	w := strings.Index(info, "wx")
	nodei := []int{}
	ip := []string{}
	wl, _ := strconv.Atoi(info[w+symbollength : w+degit+symbollength])
	fmt.Println(wl)
	for i := 0; i < wl; i++ {
		tmpint, _ := strconv.Atoi(info[symbollength+w+degit+i*degit : symbollength+w+degit+degit+i*degit])
		if i == 0 {
			nodei = append(nodei, tmpint)
		} else {
			nodei = append(nodei, nodei[i-1]+tmpint)
		}

	}

	for i := 0; i < wl; i++ {
		fmt.Println(symbollength + w + degit*wl + nodei[i])
		if i == 0 {
			ip = append(ip, info[symbollength+w+degit+degit*wl:symbollength+w+degit+degit*wl+nodei[i]])
		} else {
			ip = append(ip, info[symbollength+w+degit+degit*wl+nodei[i-1]:symbollength+w+degit+degit*wl+nodei[i]])
		}
	}
	return ip
	//fmt.Println(ip)
}

//0fa05f

func CreateNewsStation(core_arg types.CoreStruct, trans types.TransactionJson) (bool, string) {
	station := BnnTransactionToStation_new(trans)
	result := db.NewsHexGet(core_arg.Db, station.Name)
	if result.Name == "" {
		db.NewsHexPut(core_arg.Db, station.Name, station)
		return true, station.Name + " created"
	}
	return false, "name used"
}

func UpdateNewsStation(core_arg types.CoreStruct, trans types.TransactionJson) (bool, string) {
	station := BnnTransactionToStation_new(trans)
	result := db.NewsHexGet(core_arg.Db, station.Name)
	if result.Owner != crypto.KeyToAddress_hex(trans.PublicKey) {
		db.NewsHexPut(core_arg.Db, station.Name, station)
		return true, station.Name + " changed"
	}
	return false, "error user"
}
