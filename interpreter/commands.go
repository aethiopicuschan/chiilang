package interpreter

import (
	"fmt"
	"strings"
)

// エラーの定義
var ErrInvalidCode = fmt.Errorf("invalid code")

// 命令セットの定義
type Command struct {
	Chii     string
	ChiiHalf string
	BF       byte
}

// 実際の命令セットの対応
var commands = []Command{
	{"ワァ", "ﾜｧ", '>'},
	{"ワワ", "ﾜﾜ", '<'},
	{"テメ", "ﾃﾒ", '+'},
	{"ヤンパパ", "ﾔﾝﾊﾟﾊﾟ", '-'},
	{"ンショ", "ﾝｼｮ", '.'},
	{"イヤッ", "ｲﾔｯ", ','},
	{"チャル", "ﾁｬﾙ", '['},
	{"ジクジル", "ｼﾞｸｼﾞﾙ", ']'},
}

// 命令セットの情報を取得する
func GetCommands() []Command {
	return commands
}

// 独自の命令セットからなるコードが正しいかどうかを判定する関数
// ただし、ループの対応などのエラー検知は行わない
func IsValidAsChiiLang(code string) bool {
	for _, command := range commands {
		code = strings.ReplaceAll(code, command.Chii, "")
		code = strings.ReplaceAll(code, command.ChiiHalf, "")
	}
	return code == ""
}

// Brainfuckの命令セットからなるコードが正しいかどうかを判定する関数
// ただし、ループの対応などのエラー検知は行わない
func IsValidAsBrainfuck(code string) bool {
	for _, command := range commands {
		code = strings.ReplaceAll(code, string(command.BF), "")
	}
	return code == ""
}

// 独自の命令セットからなるコードをBrainfuckに変換する関数
func Convert(code string) (bf string, err error) {
	if !IsValidAsChiiLang(code) {
		err = ErrInvalidCode
		return
	}
	for _, command := range commands {
		code = strings.ReplaceAll(code, command.Chii, string(command.BF))
		code = strings.ReplaceAll(code, command.ChiiHalf, string(command.BF))
	}
	bf = code
	return
}

// Brainfuckを独自の命令セットからなるコードに変換する関数
// halfがtrueの場合は半角文字に変換する
func ConvertBack(code string, half bool) (cl string, err error) {
	if !IsValidAsBrainfuck(code) {
		err = ErrInvalidCode
		return
	}
	for _, command := range commands {
		if half {
			code = strings.ReplaceAll(code, string(command.BF), command.ChiiHalf)
		} else {
			code = strings.ReplaceAll(code, string(command.BF), command.Chii)
		}
	}
	cl = code
	return
}
