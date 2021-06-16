package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func main() {
	var cmd = &cobra.Command{
		Short:     "hello",
		Args:      cobra.OnlyValidArgs,
		ValidArgs: []string{"always", "never", "auto"},
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello, World!")
		},
	}
	cmd.Execute()
}
