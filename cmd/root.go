/*
Copyright © 2023  Matt Morrison
*/
package cmd

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

var input string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "strcase",
	Short: "String case conversion utility.",
	Long: `CLI utility to convert strings from one case to another. Supported cases include:
* lower
* upper
* snake
* kebab
* camel
`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)
		for {
			in, err := reader.ReadString('\n')
			if err != nil {
				if err != io.EOF {
					fmt.Fprintf(os.Stderr, "Error reading stdin: %v\n", err)
				}
				break
			}
			input = in
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
