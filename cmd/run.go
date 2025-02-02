package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/aethiopicuschan/chiilang/interpreter"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run [foo.chii]",
	Short: "Run a chii file",
	Args:  cobra.ExactArgs(1),
	RunE:  run,
}

func init() {
	rootCmd.AddCommand(runCmd)
}

func run(cmd *cobra.Command, args []string) (err error) {
	// 指定されたファイルを読み込む
	chii := args[0]
	if !strings.HasSuffix(chii, ".chii") {
		return fmt.Errorf("invalid file extension")
	}
	b, err := os.ReadFile(chii)
	if err != nil {
		return
	}

	// 行で区切る
	code := string(b)
	codes := strings.Split(code, "\n")

	// インタプリタを作成して実行する
	i := interpreter.NewInterpreter(30000, interpreter.Convert)
	for _, c := range codes {
		if c == "" {
			continue
		}
		output, err := i.Run(c, "")
		if err != nil {
			return err
		}
		fmt.Println(output)
	}

	return
}
