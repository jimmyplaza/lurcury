package db

import (
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/syndtr/goleveldb/leveldb"
	//"lurcury/core/block"
	//"lurcury/core/transaction"
	"lurcury/types"
)

type mleveldb struct {
	leveldb leveldb.DB
}

func OpenDB(path string) *leveldb.DB {
	db, err := leveldb.OpenFile(path, nil)
	if err != nil {
		fmt.Println(err)
	}
	return db
}

func /*(db *mleveldb)*/ Get(db *leveldb.DB, key []byte) []byte {
	data, _ := db.Get(key, nil)
	return data
}

func /*(db *mleveldb)*/ Put(db *leveldb.DB, key []byte, data []byte) {
	db.Put(key, data, nil)
}

func /*(db *mleveldb)*/ Delete(db *leveldb.DB, key []byte) {
	db.Delete(key, nil) //, nil)
}

func (db *mleveldb) FaceHexPut(keys string, data interface{}) {
	key, _ := hex.DecodeString(keys)
	data_byte, _ := json.Marshal(data)
	db.leveldb.Put(key, data_byte, nil)
}

func (db *mleveldb) FaceHexGet(keys string, target interface{}) {
	key, _ := hex.DecodeString(keys)
	data, _ := db.leveldb.Get(key, nil)
	json.Unmarshal(data, &target)
}

func FaceHexPut(db *leveldb.DB, keys string, data interface{}) {
	key, _ := hex.DecodeString(keys)
	data_byte, _ := json.Marshal(data)
	db.Put(key, data_byte, nil)
}

func FaceHexGet(db *leveldb.DB, keys string, target interface{}) {
	key, _ := hex.DecodeString(keys)
	data, _ := db.Get(key, nil)
	//var inter interface{}
	json.Unmarshal(data, &target)
}

func /*(db *mleveldb)*/ AccountHexPut(db *leveldb.DB, keys string, data types.AccountData) {
	key, _ := hex.DecodeString(keys)
	if len(key) == 0 {
		fmt.Printf("[AccountHexPut] Hex DecodeString keys is null!")
		return
	}
	data_byte, _ := json.Marshal(data)
	db.Put(key, data_byte, nil)
}

func /*(db *mleveldb)*/ AccountHexGet(db *leveldb.DB, keys string) types.AccountData {
	// k := fmt.Sprintf("\n[AccountHexGet] ===============keys: [%s]", keys)
	// fmt.Println(k)
	key, _ := hex.DecodeString(keys)
	// fmt.Println("key: ", key)
	// fmt.Println("len(key): ", len(key))
	// if len(key) == 0 {
	// 	fmt.Println("[AccountHexGet] Hex DecodeString keys is null!")
	// 	return (types.AccountData{})
	// }
	data, err := db.Get(key, nil)
	if err != nil {
		fmt.Printf("[AccountHexGet] keys[%s] err: %s \n", keys, err)
		// return (types.AccountData{})
	}
	// fmt.Println("\n[AccountHexGet] key value: ", string(data))
	inter := types.AccountData{}
	json.Unmarshal(data, &inter)
	return inter
}

func /*(db *mleveldb)*/ BlockNumberPut(db *leveldb.DB, keys string, data string) {
	//key,_ := hex.DecodeString(keys)
	//data_byte, _ := json.Marshal(data)
	db.Put([]byte(keys), []byte(data), nil)
}

func /*(db *mleveldb)*/ BlockNumberGet(db *leveldb.DB, keys string) string {
	//key,_ := hex.DecodeString(keys)
	data, _ := db.Get([]byte(keys), nil)
	//inter := types.BlockJson{}
	//json.Unmarshal(data, &inter)
	return string(data)
}

func /*(db *mleveldb)*/ BlockHexPut(db *leveldb.DB, keys string, data types.BlockJson) {
	key, _ := hex.DecodeString(keys)
	data_byte, _ := json.Marshal(data)
	db.Put(key, data_byte, nil)
}

func /*(db *mleveldb)*/ BlockHexGet(db *leveldb.DB, keys string) types.BlockJson {
	key, _ := hex.DecodeString(keys)
	data, _ := db.Get(key, nil)
	inter := types.BlockJson{}
	json.Unmarshal(data, &inter)
	return inter
}

func /*(db *mleveldb)*/ HexKeyPut(db *leveldb.DB, keys string, data string) {
	key, _ := hex.DecodeString(keys)
	//data_byte, _ := json.Marshal(data)
	db.Put(key, []byte(data), nil)
}

func /*(db *mleveldb)*/ HexKeyGet(db *leveldb.DB, keys string) string {
	key, _ := hex.DecodeString(keys)
	data, _ := db.Get(key, nil)
	return string(data)
}

func /*(db *mleveldb)*/ BlockPut(db *leveldb.DB, key []byte, data types.BlockJson) {
	data_byte, _ := json.Marshal(data)
	db.Put(key, data_byte, nil)
}

func /*(db *mleveldb)*/ BlockGet(db *leveldb.DB, key []byte) types.BlockJson {
	data, _ := db.Get(key, nil)
	inter := types.BlockJson{}
	json.Unmarshal(data, &inter)
	return inter
}

func StringHexPut(db *leveldb.DB, keys string, data string) {
	//key,_ := hex.DecodeString(keys)
	//data_byte, _ := json.Marshal(data)
	key := []byte(keys)
	db.Put(key, []byte(data), nil)
}

func StringHexGet(db *leveldb.DB, keys string) string {
	//key,_ := hex.DecodeString(keys)
	key := []byte(keys)
	data, _ := db.Get(key, nil)
	return string(data)
}

func /*(db *mleveldb)*/ TransactionPut(db *leveldb.DB, key []byte, data types.TransactionJson) {
	data_byte, _ := json.Marshal(data)
	db.Put(key, data_byte, nil)
}

func /*(db *mleveldb)*/ TransactionGet(db *leveldb.DB, key []byte) types.TransactionJson {
	data, _ := db.Get(key, nil)
	inter := types.TransactionJson{}
	json.Unmarshal(data, &inter)
	return inter
}

func /*(db *mleveldb)*/ TransactionHexPut(db *leveldb.DB, keys string, data types.TransactionJson) {
	key, _ := hex.DecodeString(keys)
	data_byte, _ := json.Marshal(data)
	db.Put(key, data_byte, nil)
}

func /*(db *mleveldb)*/ TransactionHexGet(db *leveldb.DB, keys string) types.TransactionJson {
	key, _ := hex.DecodeString(keys)
	data, _ := db.Get(key, nil)
	inter := types.TransactionJson{}
	json.Unmarshal(data, &inter)
	return inter
}

func /*(db *mleveldb)*/ NewsHexPut(db *leveldb.DB, keys string, data types.NewsStation) {
	key, _ := hex.DecodeString(keys)
	data_byte, _ := json.Marshal(data)
	db.Put(key, data_byte, nil)
}

func /*(db *mleveldb)*/ NewsHexGet(db *leveldb.DB, keys string) types.NewsStation {
	key, _ := hex.DecodeString(keys)
	data, _ := db.Get(key, nil)
	inter := types.NewsStation{}
	json.Unmarshal(data, &inter)
	return inter
}
