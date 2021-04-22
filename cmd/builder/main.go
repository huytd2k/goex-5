package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/huytd2k/site-builder/pkg/link"
)

func main() {

	urlFlag := flag.String("url", "https://gophercises.com", "the url to build site map")
	fmt.Println(*urlFlag)
	resp, err := http.Get(*urlFlag)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	// body, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	links, err := link.Parse(resp.Body)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", links)
}
