package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "greeting-app", Short: "A CLI application for greeting the user"}

	var name string

	var greetCmd = &cobra.Command{
		Use:   "greet",
		Short: "Print a greeting message",
		Run: func(cmd *cobra.Command, args []string) {
			if name == "" {
				fmt.Println("Hello there!")
			} else {
				fmt.Printf("Hello, %s!\n", name)
			}
		},
	}

	greetCmd.Flags().StringVarP(&name, "name", "n", "", "Specify a name for the greeting")

	rootCmd.AddCommand(greetCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
