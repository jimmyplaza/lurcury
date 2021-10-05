package main
import (

        "fmt"
        "net/http"
        "io/ioutil"

)

func main(){
        re , _ := httpGet("http://192.168.51.203:9003/getBlockNum")
        re2 , _ := httpGet("http://192.168.51.203:9003/getBlockbyID?blockID="+re)
        fmt.Println(re2)
}


func httpGet(url string)(string,error) {
        resp, err := http.Get(url)
        if err != nil {
                return "fail",err
        }

        defer resp.Body.Close()
        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
                return "fail",err
        }

        return string(body),err
}
