package cmd

import (
	"fmt"
	"log"

	"github.com/cmurphy/gophercises/task/db"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all of your incomplete tasks",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("You have the following tasks:")
		tasks, err := db.All()
		if err != nil {
			log.Fatal(err)
		}
		for i, t := range tasks {
			fmt.Printf("%d. %s\n", i+1, t.Value)
		}
	},
}
