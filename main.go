package main

import (
	//"log"
	"fmt"
	"flag"
	//"strings"
	"encoding/json"
	"net/http"
	"io/ioutil"
	//"github.com/jmcvetta/napping"
)

// JSONデコード用構造体
type Result struct {
	Name string `json:"name"`
	Tld []string `json:"topLevelDomain"`
	A2c string `json:"alpha2Code"`
	A3c string `json:"alpha3Code"`
}

func main() {
	flag.Parse()
	url := "https://restcountries.eu/rest/v2/alpha/"
	url += flag.Arg(0)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	var result Result
	if err := json.Unmarshal(bytes, &result); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result.Name)
}
