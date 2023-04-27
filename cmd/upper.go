/*
Copyright © 2023 Matt Morrison
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// upperCmd represents the upper command
var upperCmd = &cobra.Command{
	Use:   "upper",
	Short: "Convert the input to uppercase",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(strings.ToUpper(input))
	},
}

func init() {
	rootCmd.AddCommand(upperCmd)
}
