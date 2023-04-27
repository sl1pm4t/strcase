/*
Copyright © 2023 Matt Morrison
*/
package cmd

import (
	"fmt"
	"github.com/stoewer/go-strcase"

	"github.com/spf13/cobra"
)

// camelCmd represents the camel command
var camelCmd = &cobra.Command{
	Use:   "camel",
	Short: "Convert the input to CamelCase",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(strcase.UpperCamelCase(input))
	},
}

func init() {
	rootCmd.AddCommand(camelCmd)
}
