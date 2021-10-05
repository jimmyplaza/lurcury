package types

import()

type Content struct{
        Title string `json:"title"`
        Intro string `json:"intro"`
        Time string `json:"time"`
	Index string `json:"index"`
}

type Nodeinfo struct{
	Ip string `json:"ip"`
	Name string `json:"ip"`
}

type NewsStation struct{
	Name string `json:"name"`
	Intro string `json:"intro"`
	//Time string `json:"time"`
	Node []string `json:"node"`
	Article []Content `json:"article"`
	Picture []Content `json:"picture"`
	Video []Content `json:"video"`
	Owner string `json:"owner"`
}


