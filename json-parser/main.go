package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	data, err := ioutil.ReadFile("C:\\Users\\Ant\\Desktop\\Sessions - 10-08-2018 06-05-44.json")
	check(err)

	var decodedData []map[string]interface{}
	err = json.Unmarshal(data, &decodedData)
	check(err)

	f, err := os.Create("C:\\Users\\Ant\\Desktop\\sessionsUrls.txt")
	check(err)
	defer f.Close()

	for _, v := range decodedData {
		for _, v1 := range v["windows"].(map[string]interface{}) {
			for _, v2 := range v1.(map[string]interface{}) {
				var t = v2.(map[string]interface{})
				f.WriteString(t["url"].(string) + "\n" + t["title"].(string) + "\n\n")
			}
		}
	}
}
