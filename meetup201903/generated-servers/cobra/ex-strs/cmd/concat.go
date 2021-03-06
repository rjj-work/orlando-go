// Copyright © 2019 Ralph J. Jackson <rjj.work@gmail.com>

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/rjj-work/orlando-go/meetup201903/generated-servers/cobra/ex-strs/concat"
)

// concatCmd represents the concat command
var concatCmd = &cobra.Command{
	Use:   "concat",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("concat called args: %v\n", args)

		s := concat.Doit(args)
		fmt.Println(s)
	},
}

func init() {
	rootCmd.AddCommand(concatCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// concatCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// concatCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
