/*
Copyright © 2024 Aayush Ranjan
*/
package cmd

import (
	// "fmt"
	"log"
	"github.com/aayushxrj/pluto/todo"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new todo",
	Long: `Add will create a new todo item to the list`,
	Run: addRun,
}

func addRun(cmd *cobra.Command, args []string){

	// items := []todo.Item{}
	// why ?
	// var items = []todo.Item{}

	items, err := todo.ReadItems(dataFile)
	if err != nil {
		log.Printf("%v", err)
	}
	// why ?

	for _,x := range args{
		// items = append(items, todo.Item{Text: x})
		item:= todo.Item{Text:x}
		item.SetPriority(priority)
		items = append(items, item)
	}
	// fmt.Printf("%#v\n", items)
	err = todo.SaveItems(dataFile, items)
	if err!=nil {
		log.Printf("%v", err)
	}
}

var priority int

func init() {
	rootCmd.AddCommand(addCmd)

	//added
	addCmd.Flags().IntVarP(&priority,
		"priority",
		"p",
		2,
		"Priority:1,2,3")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
