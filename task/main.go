package main

import (
	"github.com/cmurphy/gophercises/task/cmd"
	"github.com/cmurphy/gophercises/task/db"
)

func main() {
	db.Init()
	cmd.Execute()
}
