package main

import (
	"bufio"
	"fmt"
	"github.com/cmurphy/gophercises/link"
	"os"
)

func main() {
	file, err := os.Open("ex4.html")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	r := bufio.NewReader(file)
	links, err := link.Parse(r)
	if err != nil {
		panic(err)
	}
	for _, l := range links {
		fmt.Printf("%s\t-->\t%s\n", l.Href, l.Text)
	}
}
