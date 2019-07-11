package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func pra_1_8() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		_, err = io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
	}
}
func pra_1_9() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%s\n", resp.Status)
	}
}
func main() {
	// for _, url := range os.Args[1:] {
	// 	resp, err := http.Get(url)
	// 	if err != nil {
	// 		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
	// 		os.Exit(1)
	// 	}
	// 	// b, err := ioutil.ReadAll(resp.Body)
	// 	// resp.Body.Close()
	// 	//practice 1.7
	// 	_, err = io.Copy(os.Stdout, resp.Body)

	// 	if err != nil {
	// 		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
	// 		os.Exit(1)
	// 	}
	// 	//fmt.Printf("%s", b)
	// }
	pra_1_8()
	pra_1_9()
}
