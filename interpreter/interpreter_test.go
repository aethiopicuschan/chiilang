package interpreter_test

import (
	"testing"

	"github.com/aethiopicuschan/chiilang/interpreter"
	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	testcases := []struct {
		name        string
		code        string
		input       string
		expected    string
		expectedErr error
	}{
		{
			name:        "hello world",
			code:        "テメテメテメテメテメテメテメテメテメチャルワァテメテメテメテメテメテメテメテメワァテメテメテメテメテメテメテメテメテメテメテメワァテメテメテメテメテメワワワワワワヤンパパジクジルワァンショワァテメテメンショテメテメテメテメテメテメテメンションショテメテメテメンショワァヤンパパンショヤンパパヤンパパヤンパパヤンパパヤンパパヤンパパヤンパパヤンパパヤンパパヤンパパヤンパパヤンパパンショワワテメテメテメテメテメテメテメテメンショヤンパパヤンパパヤンパパヤンパパヤンパパヤンパパヤンパパヤンパパンショテメテメテメンショヤンパパヤンパパヤンパパヤンパパヤンパパヤンパパンショヤンパパヤンパパヤンパパヤンパパヤンパパヤンパパヤンパパヤンパパンショワァテメンショ",
			input:       "",
			expected:    "Hello, world!",
			expectedErr: nil,
		},
		{
			name:        "echo",
			code:        "ｲﾔｯﾁｬﾙﾝｼｮﾁｬﾙﾔﾝﾊﾟﾊﾟｼﾞｸｼﾞﾙｲﾔｯｼﾞｸｼﾞﾙ",
			input:       "橋本",
			expected:    "橋本",
			expectedErr: nil,
		},
		{
			name:        "invalid code",
			code:        "ﾊﾁﾃﾒ",
			input:       "",
			expected:    "",
			expectedErr: interpreter.ErrInvalidCode,
		},
		{
			name:        "unmatched bracket",
			code:        "ﾁｬﾙﾝｼｮﾁｬﾙﾔﾝﾊﾟﾊﾟｼﾞｸｼﾞﾙ",
			input:       "",
			expected:    "",
			expectedErr: interpreter.ErrUnmatchedBracket,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			i := interpreter.NewInterpreter(30000, interpreter.Convert)
			actual, err := i.Run(tc.code, tc.input)
			require.Equal(t, tc.expected, actual)
			require.Equal(t, tc.expectedErr, err)
		})
	}

}
