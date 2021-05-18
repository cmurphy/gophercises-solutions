package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/cmurphy/gophercises/task/db"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task to your TODO list",
	Run: func(cmd *cobra.Command, args []string) {
		item := strings.Join(args, " ")
		if err := db.Add(item); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Added \"%s\" to your task list.\n", item)
	},
}
