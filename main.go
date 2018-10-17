// * * * * * * * * * * * * * *
// * * * * * * * * * * * * * *

package main

import (
	"io/ioutil"
	"net/http"

	"encoding/json"
	"fmt"
	"log"
	"os"
)

// Result data type
type Result struct {
	Identifier string  `json:"identifier"`
	Item       string  `json:"item"`
	Value      float64 `json:"value"`
}

// Configuration pulled from json file
type Configuration struct {
	Username string
	Password string
	Url string
}

func loadConfiguration() Configuration {
	var configuration Configuration
	file, err := os.Open("./.config.json")
	if err != nil {
		log.Fatal(err)
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&configuration)
	if err != nil {
		log.Fatal(err)
	}
	return configuration
}

//
func getDataPoint(identifier string, dataPoint string) Result {
	conf := loadConfiguration()
	client := &http.Client{}
	url := conf.Url + "/data_point?identifier=" + identifier + "&" + "item=" + dataPoint
	req, err := http.NewRequest("GET", url, nil)
	if err != nil{
		log.Fatal(err)
	}
	req.SetBasicAuth(conf.Username, conf.Password)
	req.Header.Set("User-Agent", "Spider_Bot/1.0")
	resp, err := client.Do(req)
	if err != nil{
		log.Fatal(err)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	var ret Result
	err = json.Unmarshal(bodyText, &ret)
	if err != nil{
		log.Fatal(err)
	}
	return ret
}

func main() {
	args := os.Args
	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "Usage: %s\n", args[0])
		os.Exit(1)
	}
	result := getDataPoint("AAPL", "ask_price")
	fmt.Println(result)
}
