package main

import(
        "io/ioutil"
        "net/http"
)

func Get(url string)([]byte, error){
        resp, err := http.Get(url)
        if(err != nil){
                return []byte(""),err
        }else{
                body, err := ioutil.ReadAll(resp.Body)
                if(err != nil){
                        return []byte(""),err
                }else{
                        return body, err
                }
        }
}
