package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "mycli"}

	var cmdHello = &cobra.Command{
		Use:   "hello",
		Short: "Affiche un message",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello Kubernetes!")
		},
	}

	rootCmd.AddCommand(cmdHello)
	rootCmd.Execute()
}
