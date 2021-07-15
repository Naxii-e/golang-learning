package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	//Printfはフォーマットを含む
	//Printlnはフォーマットを含まない 改行あり
	//Printはフォーマットを含まない　改行なし

	url := "https://api.mcsrvstat.us/2/hypixel.net"
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	client := new(http.Client)
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Error Request: ", err)
		return
	}

	defer resp.Body.Close() //TCPコネクションをクローズする

	if resp.StatusCode != 200 {
		fmt.Println("Error Response:", resp.Status)
		return
	}

	fmt.Printf("%-v", resp)

	body, _ := ioutil.ReadAll(resp.Body)

	var Server Root
	json.Unmarshal(body, Server)
	fmt.Printf("結果: %-v", Server)

}

//WebAPIから返されるjsonの構造体を準備
type Root struct {
	Value []Server `json:"value"`
}

type Server struct {
	ip   string
	port int
	ping bool
}
