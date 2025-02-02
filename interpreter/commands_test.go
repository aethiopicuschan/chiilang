package interpreter_test

import (
	"testing"

	"github.com/aethiopicuschan/chiilang/interpreter"
	"github.com/stretchr/testify/require"
)

func TestIsValidAsChiiLang(t *testing.T) {
	testcases := []struct {
		name     string
		code     string
		expected bool
	}{
		{
			name:     "valid code",
			code:     "ワァワワテメヤンパパンショイヤッチャルジクジル",
			expected: true,
		},
		{
			name:     "valid code with half-width characters",
			code:     "ﾜｧﾜﾜﾃﾒﾔﾝﾊﾟﾊﾟﾝｼｮｲﾔｯﾁｬﾙｼﾞｸｼﾞﾙ",
			expected: true,
		},
		{
			name:     "invalid code",
			code:     "ﾊﾁﾃﾒ",
			expected: false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			actual := interpreter.IsValidAsChiiLang(tc.code)
			require.Equal(t, tc.expected, actual)
		})
	}
}

func TestIsValidAsBrainfuck(t *testing.T) {
	testcases := []struct {
		name     string
		code     string
		expected bool
	}{
		{
			name:     "valid code",
			code:     "><+-.,[]",
			expected: true,
		},
		{
			name:     "invalid code",
			code:     "><+-.,[]a",
			expected: false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			actual := interpreter.IsValidAsBrainfuck(tc.code)
			require.Equal(t, tc.expected, actual)
		})
	}
}

func TestConvert(t *testing.T) {
	testcases := []struct {
		name        string
		code        string
		expected    string
		expectedErr error
	}{
		{
			name:        "convert to Brainfuck",
			code:        "ワァワワテメヤンパパンショイヤッチャルジクジル",
			expected:    "><+-.,[]",
			expectedErr: nil,
		},
		{
			name:        "convert to Brainfuck with half-width characters",
			code:        "ﾜｧﾜﾜﾃﾒﾔﾝﾊﾟﾊﾟﾝｼｮｲﾔｯﾁｬﾙｼﾞｸｼﾞﾙ",
			expected:    "><+-.,[]",
			expectedErr: nil,
		},
		{
			name:        "invalid code",
			code:        "ﾊﾁﾃﾒ",
			expected:    "",
			expectedErr: interpreter.ErrInvalidCode,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			actual, err := interpreter.Convert(tc.code)
			require.Equal(t, tc.expected, actual)
			require.Equal(t, tc.expectedErr, err)
		})
	}
}

func TestConvertBack(t *testing.T) {
	testcases := []struct {
		name        string
		code        string
		half        bool
		expected    string
		expectedErr error
	}{
		{
			name:        "convert to ChiiLang",
			code:        "><+-.,[]",
			half:        false,
			expected:    "ワァワワテメヤンパパンショイヤッチャルジクジル",
			expectedErr: nil,
		},
		{
			name:        "convert to ChiiLang with half-width characters",
			code:        "><+-.,[]",
			half:        true,
			expected:    "ﾜｧﾜﾜﾃﾒﾔﾝﾊﾟﾊﾟﾝｼｮｲﾔｯﾁｬﾙｼﾞｸｼﾞﾙ",
			expectedErr: nil,
		},
		{
			name:        "invalid code",
			code:        "><+-.,[]a",
			half:        false,
			expected:    "",
			expectedErr: interpreter.ErrInvalidCode,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			actual, err := interpreter.ConvertBack(tc.code, tc.half)
			require.Equal(t, tc.expected, actual)
			require.Equal(t, tc.expectedErr, err)
		})
	}
}
