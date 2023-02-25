package main

import (
	"encoding/json"
	"fmt"

	"github.com/bitly/go-simplejson"
)

func main() {
	// 解析json
	var s ServerSlice
	str := `{"Server":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
	json.Unmarshal([]byte(str), &s)
	fmt.Println(s)

	// 生成json
	s.Server = append(s.Server, Server{ServerName: "Shanghai_VPN", ServerIp: "127.0.0.1"})
	s.Server = append(s.Server, Server{ServerName: "Beijing_VPN", ServerIp: "127.0.0.2"})
	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))

	// simplejson 解析未知json
	json, err := simplejson.NewJson([]byte(`{
		"test": {
			"array": [1, "2", 3],
			"int": 10,
			"float": 5.150,
			"bignum": 9223372036854775807,
			"string": "simplejson",
			"bool": true
		}
	}`))

	if err != nil {

	}

	array, _ := json.Get("test").Get("array").Array()
	i, _ := json.Get("test").Get("int").Int()
	fmt.Println(array)
	fmt.Println(i)

}

type Server struct {
	ServerName string
	ServerIp   string
}

type ServerSlice struct {
	Server []Server
}
