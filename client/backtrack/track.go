package backtrack

import (
        "encoding/json"
	//"errors"
        "fmt"
        "io/ioutil"
        "lurcury/db"
        "lurcury/types"
        "net/http"
)
func GET(url string)([]byte, error) {
    client := &http.Client{}
    reqest, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return []byte(""),err
    }
    reqest.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
    reqest.Header.Add("Accept-Language", "ja,zh-CN;q=0.8,zh;q=0.6")
    reqest.Header.Add("Connection", "keep-alive")
    reqest.Header.Add("Cookie", "设置cookie")
    reqest.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:12.0) Gecko/20100101 Firefox/12.0")
    response, err := client.Do(reqest)
    defer response.Body.Close()
    cookies := response.Cookies()
    for _, cookie := range cookies {
        fmt.Println("cookie:", cookie)
    }

    body, err1 := ioutil.ReadAll(response.Body)
    if err1 != nil {
        return []byte(""),err1
    }
    return body, nil
}
func GetUpdateAccount(core_arg *types.CoreStruct, target string)(string) {
	b, err := GET(core_arg.Config.Peers[0]+"/getAccount?address="+target)
	var obj types.AccountData
	if err == nil{
		if err := json.Unmarshal(b, &obj); err != nil {
			fmt.Println(err)
		}
		db.AccountHexPut(core_arg.Db, obj.Address, obj)
		return "success"
	}
	return "error"
}

