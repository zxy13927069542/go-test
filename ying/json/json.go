package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/**
顶层结点
*/
type Response struct {
	Status  int      `json: "status"`
	Message string   `json: "message"`
	Result  []Result `json:"result"`
}

type Result struct {
	Name     string   `json:"name"`
	Location Location `json:"location"`
	Uid      string   `json:"uid"`
	Province string   `json:"province"`
	City     string   `json:"city"`
	District string   `json:"district"`
	Business string   `json:"business"`
	Cityid   string   `json:"cityid"`
	Tag      string   `json:"tag"`
	Address  string   `json:"address"`
	Children []Child  `json:"children"`
	Adcode   string   `json:"adcode"`
}

type Location struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type Child struct {
	Uid       string `json:"uid"`
	Name      string `json:"name"`
	Show_name string `json:"show_name"`
}

/*type People struct {
	name	string		`json:"name"`
	craft	string		`json:"craft"`
}

type Response struct {
	message string		`json:"message"`
	people	[]People	`json:"people"`
	number	int			`json:"number"`
}*/

func main() {
	url := "https://api.map.baidu.com/place/v2/suggestion?query=天安门&region=北京&city_limit=true&output=json&ak=pqnjDWUVvNlGj7vVdxhzD6zLoi0Df1bG"

	//发送get请求
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	//解码
	var result Response
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%#v\n", result)
}
