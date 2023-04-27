/*
Copyright © 2023 Matt Morrison
*/
package cmd

import (
	"fmt"
	"github.com/stoewer/go-strcase"

	"github.com/spf13/cobra"
)

// kebabCmd represents the kebab command
var kebabCmd = &cobra.Command{
	Use:   "kebab",
	Short: "Convert the input to kebab-case",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(strcase.KebabCase(input))
	},
}

func init() {
	rootCmd.AddCommand(kebabCmd)
}
