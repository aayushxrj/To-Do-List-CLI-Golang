/*
Copyright © 2024 Aayush Ranjan
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"sort"

	// "strconv"

	// "strconv"
	"text/tabwriter"

	"github.com/aayushxrj/pluto/todo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the todos",
	Long: `Listing the todos`,
	Run: listRun,
}

// lists all the todos
func listRun(cmd *cobra.Command, args []string){
	items, err := todo.ReadItems(viper.GetString("dataFile"))
	if err!= nil{
		log.Printf("%v", err)
	}

	//sort
	sort.Sort(todo.ByPri(items))

	w := tabwriter.NewWriter(os.Stdout, 3, 0, 1, ' ', 0)
	// for _, i := range items {
	// 	fmt.Fprintln(w,i.Label() +"\t"+i.PrettyDone() +"\t"+ i.PrettyP() +"\t"+ i.Text +"\t")
	// 	// fmt.Fprintln(w, strconv.Itoa(priority) +"\t"+ i.Text +"\t")
	// }
	for _, i := range items {
		if(allOpt || i.Done == doneOpt){
			fmt.Fprintln(w,i.Label() +"\t"+i.PrettyDone() +"\t"+ i.PrettyP() +"\t"+ i.Text +"\t")
		}
	}
	w.Flush()
}

var (
	doneOpt bool
	allOpt bool
)


func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolVar(&doneOpt,
	"done",
	false,
	"Show 'Done' Todos")
	listCmd.Flags().BoolVar(&allOpt,
		"all",
		false,
		"Show all Todos")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
