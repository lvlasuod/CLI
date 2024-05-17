/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"sort"
	"text/tabwriter"

	"github.com/lvlasuod/timo/todo"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the todos",
	Long:  `Listing the todos`,
	Run:   listRun,
}

var (
	doneOpt bool
	allOpt  bool
)

func listRun(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(dataFile)
	if err != nil {

		log.Printf("%v", err)
	}

	sort.Sort(todo.ByPriority(items))

	w := tabwriter.NewWriter(os.Stdout, 3, 0, 1, ' ', 0)
	fmt.Fprintln(w, "====="+"\t"+"========"+"\t"+"===="+"\t"+"===="+"\t")
	fmt.Fprintln(w, "Label"+"\t"+"Priority"+"\t"+"Item"+"\t"+"Done"+"\t")
	fmt.Fprintln(w, "====="+"\t"+"========"+"\t"+"===="+"\t"+"===="+"\t")
	for _, i := range items {
		if allOpt || i.Done == doneOpt {
			fmt.Fprintln(w, i.Label()+"\t"+i.PrettyP()+"\t"+i.Text+"\t"+i.PrettyDone()+"\t")
		}
	}
	w.Flush()
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	listCmd.Flags().BoolVarP(&doneOpt, "done", "d", false, "Show 'Done' todos")
	listCmd.Flags().BoolVarP(&allOpt, "all", "a", false, "Show all todos")

}
