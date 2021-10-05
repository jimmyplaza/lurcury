package transaction

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func Box(arg string)(string, error){
    url := "http://127.0.0.1:5314/casigo/box"

    var jsonStr = []byte(arg)
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return "http err", err
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil{
        return "ioutil err", err
    }
    return string(body), nil
}

func Excute(arg string)(string, error){
    url := "http://127.0.0.1:5314/casigo/excute"
    var jsonStr = []byte(arg)
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return "http err", err
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil{
        return "ioutil err", err
    }
    return string(body), nil
}

func Call(arg string)(string, error){
    url := "http://127.0.0.1:5314/casigo/call"

    var jsonStr = []byte(`{"func": "b(444);", "account":"var sender = '5';", "address": "0x"}`)
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return "http err", err
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil{
        return "ioutil err", err
    }
    return string(body), nil

}

func Deploy(arg string)(string,error){
    url := "http://127.0.0.1:5314/casigo/deploy"

    //var jsonStr = []byte(`{"init": "var x=0; var y=1;", "codes":"function b(input){x += 40; y = 17; var d = 7; x=input}", "account":"var sender = '5';","address": "0x"}`)
    var jsonStr = []byte(arg)

    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
	return "http err", err
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil{
        return "ioutil err", err
    }
    return string(body), nil

}
