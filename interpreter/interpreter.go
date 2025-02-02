package interpreter

import (
	"fmt"
)

var ErrUnmatchedBracket = fmt.Errorf("unmatched bracket")

// インタプリタ
type Interpreter struct {
	memory    []byte
	pointer   int
	loopPairs map[int]int // ループの対応関係を記録
	converter func(string) (string, error)
}

// メモリサイズを指定してインタプリタを生成する
func NewInterpreter(memorySize int, converter func(string) (string, error)) *Interpreter {
	return &Interpreter{
		memory:    make([]byte, memorySize),
		pointer:   0,
		loopPairs: make(map[int]int),
		converter: converter,
	}
}

// ループのペアを解析する関数
func (i *Interpreter) preprocessLoops(code string) (err error) {
	stack := []int{}
	for index, command := range code {
		if command == '[' {
			stack = append(stack, index)
		} else if command == ']' {
			if len(stack) == 0 {
				err = ErrUnmatchedBracket
				return
			}
			open := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			i.loopPairs[open] = index
			i.loopPairs[index] = open
		}
	}
	if len(stack) > 0 {
		err = ErrUnmatchedBracket
	}
	return
}

// コードを実行する関数
func (i *Interpreter) Run(code string, input string) (output string, err error) {
	// 独自の命令セットをBrainfuckに変換
	code, err = i.converter(code)
	if err != nil {
		return
	}

	// UTF-8 をバイト列として処理
	inputBytes := []byte(input)
	var outputBytes []byte

	// ループのペアを解析
	err = i.preprocessLoops(code)
	if err != nil {
		return "", err
	}

	si := 0
	inputIndex := 0

	for si < len(code) {
		switch code[si] {
		case '+':
			i.memory[i.pointer]++
		case '-':
			i.memory[i.pointer]--
		case '>':
			i.pointer++
		case '<':
			i.pointer--
		case ',':
			if inputIndex < len(inputBytes) {
				i.memory[i.pointer] = inputBytes[inputIndex]
				inputIndex++
			} else {
				i.memory[i.pointer] = 0
			}
		case '.':
			outputBytes = append(outputBytes, i.memory[i.pointer])
		case '[':
			if i.memory[i.pointer] == 0 {
				si = i.loopPairs[si]
			}
		case ']':
			if i.memory[i.pointer] != 0 {
				si = i.loopPairs[si]
			}
		}
		si++
	}

	output = string(outputBytes)
	return
}
