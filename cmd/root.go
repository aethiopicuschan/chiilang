package cmd

import (
	"fmt"
	"runtime/debug"

	"github.com/aethiopicuschan/chiilang/interpreter"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:  "chiilang",
	Long: "chiilang is a small and cute programming language",
	RunE: root,
}

func init() {
	bi, ok := debug.ReadBuildInfo()
	if ok {
		rootCmd.Version = bi.Main.Version
	}
}

func root(cmd *cobra.Command, args []string) (err error) {
	i := interpreter.NewInterpreter(30000, interpreter.Convert)
	for {
		fmt.Print("> ")
		var code string
		fmt.Scanln(&code)
		if code == "exit" || code == "quit" {
			break
		}
		output, err := i.Run(code, "")
		if err != nil {
			return err
		}
		fmt.Println(output)
	}
	return nil
}

func Execute() error {
	return rootCmd.Execute()
}
