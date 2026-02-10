package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type catfactresponseStruct struct {
	Fact   string `json:"fact"`
	Length int    `json:"length"`
}

func main() {

	url := "https://catfact.ninja/fact"

	resp, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	bodtBytes, err := io.ReadAll(resp.Body)

	var data catfactresponseStruct

	if err := json.Unmarshal(bodtBytes, &data); err != nil {
		fmt.Println("json unmarshal failed")
		return
	}

	fmt.Println(data.Fact, data.Length)

}
