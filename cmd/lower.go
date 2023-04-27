/*
Copyright © 2023 Matt Morrison
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

// lowerCmd represents the lower command
var lowerCmd = &cobra.Command{
	Use:   "lower",
	Short: "Convert the input to lowercase",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(strings.ToLower(input))
	},
}

func init() {
	rootCmd.AddCommand(lowerCmd)
}
