/*
Copyright © 2023 Matt Morrison
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stoewer/go-strcase"
)

// snakeCmd represents the snake command
var snakeCmd = &cobra.Command{
	Use:     "snake",
	Aliases: []string{"s"},
	Short:   "Convert the input to snake_case",
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(strcase.SnakeCase(input))
	},
}

func init() {
	rootCmd.AddCommand(snakeCmd)
}
