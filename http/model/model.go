package model

import (
	"io/ioutil"
	"encoding/json"
	"net/http"
)

//curl -X GET "localhost:9000/testparams?key=value1"
func TestPutDb(res http.ResponseWriter, req *http.Request){
        res.Header().Add("Access-Control-Allow-Origin","*")
        val := req.FormValue("key")
        res.Write([]byte(val))
}

func TestGetDb(res http.ResponseWriter, req *http.Request){
        res.Header().Add("Access-Control-Allow-Origin","*")
        val := req.FormValue("key")
        res.Write([]byte(val))
}

//curl -X GET "localhost:9000/testparams?key=value1"
func TestParams(res http.ResponseWriter, req *http.Request){
	res.Header().Add("Access-Control-Allow-Origin","*")
	val := req.FormValue("key")
	res.Write([]byte(val))
}

type TestBodys_Struct struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

//curl "localhost:9000/testbodys?key=value1" -d '{"Id":1,"Name":"123"}'
func TestBodys(res http.ResponseWriter, req *http.Request){
	res.Header().Add("Access-Control-Allow-Origin","*")
	b, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	var msg TestBodys_Struct
	json.Unmarshal(b, &msg)
	res.Header().Set("content-type", "application/json")
	res.Write([]byte(string(msg.Name)))
}

/*
func KeyStoreProduct(res http.ResponseWriter, req *http.Request){
	res.Header().Add("Access-Control-Allow-Origin","*")
	res.Write([]byte(crypto.KeyStoreOut()))
}
*/
