package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/cmurphy/gophercises/sitemap"
)

func main() {
	siteurl := flag.String("site", "", "site URL")
	maxDepth := flag.Int("depth", -1, "maximum depth to traverse, default unlimited")
	flag.Parse()
	if *siteurl == "" {
		fmt.Println("specify site with -site")
		os.Exit(1)
	}
	sm, err := sitemap.NewSiteMap(*siteurl, *maxDepth)
	if err != nil {
		panic(err)
	}
	fmt.Println(sm)
}
