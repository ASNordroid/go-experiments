package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func req(url string, ch chan<- string) {
	//time.Sleep(100 * time.Millisecond)
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	req, err := http.NewRequest("GET", url, nil)
	check(err)
	resp, err := client.Do(req)
	if err != nil {
		ch <- fmt.Sprintf("no %s", url)
	} else {
		defer resp.Body.Close()
		ch <- fmt.Sprintf("%d - %s", resp.StatusCode, url)
	}
}

func main() {
	file, err := os.Open("D:\\Documents\\firefox_links.txt")
	check(err)

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	steps := 10000
	ch := make(chan string)
	for j := 0; j < steps; j++ {
		go req(lines[j], ch)
	}

	for j := 0; j < steps; j++ {
		msg := <-ch
		if strings.Split(msg, " - ")[0] != "200" {
			fmt.Println(msg)
		}
	}

	fmt.Println("Done!")
}
