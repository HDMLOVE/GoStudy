package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintln(os.Stderr, "fentch :%v\n", err)
			os.Exit(1)
		}
		_, err = io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Println(os.Stderr, "fentch Copy failed, %v\n", err)
			os.Exit(1)
		}
		resp.Body.Close()
		fmt.Println("Status: ", resp.Status)
	}
}
