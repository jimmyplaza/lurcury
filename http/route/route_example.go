package route

import (
	"encoding/json"
	//"fmt"
	"net/http"
	"lurcury/db"
	"lurcury/http/model"
	"lurcury/types"
)

func Router_exp(coreStruct types.CoreStruct){
	TestParams := func (res http.ResponseWriter, req *http.Request){
                res.Header().Add("Access-Control-Allow-Origin","*")
		coreStruct.Db.Put(/*coreStruct,*/ []byte("top"),[]byte("top"),nil)
		val := req.FormValue("key")
                res.Write([]byte(val))
        }

	TestGet := func(res http.ResponseWriter, req *http.Request){
                res.Header().Add("Access-Control-Allow-Origin","*")
		val := req.FormValue("key")
		re := db.Get(coreStruct.Db, []byte(val))
                res.Write(re)//[]byte(re))
	}

        TestHexGet := func(res http.ResponseWriter, req *http.Request){
                res.Header().Add("Access-Control-Allow-Origin","*")
                val := req.FormValue("key")
                //re := db.Get(coreStruct.Db,[]byte(val))
		re := db.BlockHexGet(coreStruct.Db, val)
		//fmt.Println(re)
		b, _ := json.Marshal(re)
                res.Write([]byte(string(b)))//[]byte(re))
        }

	http.HandleFunc("/testparams", TestParams)
        http.HandleFunc("/testGet", TestGet)
	//curl -X GET "localhost:9000/testHexGet?key=96b3c815348b2f176580543080703f5b015f1c2df4091e3c75deaa8c89fd6a1f"
	http.HandleFunc("/testHexGet", TestHexGet)
	http.HandleFunc("/testbodys", model.TestBodys)

}

func Test(coreStruct types.CoreStruct){
        TestParams := func (res http.ResponseWriter, req *http.Request){
                res.Header().Add("Access-Control-Allow-Origin","*")
                coreStruct.Db.Put(/*coreStruct,*/ []byte("top"),[]byte("top"),nil)
                val := req.FormValue("key")
                res.Write([]byte(val))
        }

        TestGet := func(res http.ResponseWriter, req *http.Request){
                res.Header().Add("Access-Control-Allow-Origin","*")
                re,_ := coreStruct.Db.Get([]byte("top"),nil)
                //val := req.FormValue("key")
                res.Write([]byte(re))
        }

        http.HandleFunc("/ttestparams", TestParams)
        http.HandleFunc("/ttestGet", TestGet)
        http.HandleFunc("/ttestbodys", model.TestBodys)

}
