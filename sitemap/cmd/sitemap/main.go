package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/cmurphy/gophercises/sitemap"
)

func main() {
	siteurl := flag.String("site", "", "site URL")
	flag.Parse()
	if *siteurl == "" {
		fmt.Println("specify site with -site")
		os.Exit(1)
	}
	sm, err := sitemap.NewSiteMap(*siteurl)
	if err != nil {
		panic(err)
	}
	fmt.Println(sm)
}
