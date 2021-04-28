package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"strings"

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

	reqUrl := resp.Request.URL
	baseUrl := url.URL{
		Scheme: reqUrl.Scheme,
		Host:   reqUrl.Host,
	}
	base := baseUrl.String()

	var hrefs []string

	for _, l := range links {
		switch {
		case strings.HasPrefix(l.Href, "/"):
			hrefs = append(hrefs, base+l.Href)
		case strings.HasPrefix(l.Href, "http"):
			hrefs = append(hrefs, l.Href)
		}
	}

	for _, href := range hrefs {
		println(href)
	}
}
