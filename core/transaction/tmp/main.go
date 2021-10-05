package main

import (
    "fmt"
)

type After struct {
    s []string
}

func contrivedAfter() interface{} {
    return After{[]string{"new value"}}
}

func main() {
    //b := Before{map[string]string{"some": "value"}}
    a := contrivedAfter()//.(After)
    //fmt.Println(a.m)
    fmt.Println(a.s)
}
