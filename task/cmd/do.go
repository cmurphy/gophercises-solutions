package cmd

import (
	"encoding/binary"
	"fmt"
	"log"
	"strconv"

	"github.com/cmurphy/gophercises/task/db"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(doCmd)
}

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark a task on your TODO list as complete",
	Run: func(cmd *cobra.Command, args []string) {
		taskID, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf("invalid argument: %s, expected integer\n", args[0])
			return
		}
		tasks, err := db.All()
		if err != nil {
			log.Fatal(err)
		}
		task := tasks[taskID-1]
		err = db.Delete(task.Key)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("You have completed the \"%s\" task.\n", task.Value)
	},
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
