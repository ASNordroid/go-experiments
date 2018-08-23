package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	data, err := ioutil.ReadFile("C:\\Users\\Ant\\Desktop\\Sessions - 10-08-2018 06-05-44.json")
	if err != nil {
		panic(err)
	}

	var decodedData []map[string]interface{}
	if err := json.Unmarshal(data, &decodedData); err != nil {
		panic(err)
	}

	f, err := os.Create("sessionsUrls.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	defer w.Flush()

	for _, v := range decodedData {
		for _, v2 := range v["windows"].(map[string]interface{}) {
			for _, v3 := range v2.(map[string]interface{}) {
				var t = v3.(map[string]interface{})
				fmt.Fprint(w, t["url"], "\n")
				fmt.Fprint(w, t["title"], "\n\n")
			}
		}
	}
}
