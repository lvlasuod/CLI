/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"text/tabwriter"

	"github.com/lvlasuod/timo/todo"
	"github.com/spf13/cobra"
)

// findCmd represents the find command
var findCmd = &cobra.Command{
	Use:   "find",
	Short: "find a todo using text",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: findRun,
}

func findRun(cmd *cobra.Command, args []string) {
	items, _ := todo.ReadItems(dataFile)
	i, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalln(args[0], "is not a valid lable", err)
	}
	w := tabwriter.NewWriter(os.Stdout, 3, 0, 1, ' ', 0)

	if i > 0 && i < len(items) {
		fmt.Fprintf(w, items[i-1].Label()+"\t"+items[i-1].PrettyP()+"\t"+items[i-1].Text+"\t"+items[i-1].PrettyDone()+"\t")
		fmt.Printf("%q\n", items[i-1].Label()+" "+items[i-1].PrettyP()+" "+items[i-1].Text+" "+items[i-1].PrettyDone())
	} else {
		log.Println(i, "doesn't match any items")
	}

}

func init() {
	rootCmd.AddCommand(findCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// findCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// findCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
