# ChiiLang

[![License: WTFPL](https://img.shields.io/badge/License-WTFPL-brightgreen?style=flat-square)](/LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/aethiopicuschan/chiilang.svg)](https://pkg.go.dev/github.com/aethiopicuschan/chiilang)
[![Go Report Card](https://goreportcard.com/badge/github.com/aethiopicuschan/chiilang)](https://goreportcard.com/report/github.com/aethiopicuschan/chiilang)
[![CI](https://github.com/aethiopicuschan/chiilang/actions/workflows/ci.yaml/badge.svg)](https://github.com/aethiopicuschan/chiilang/actions/workflows/ci.yaml)

`ChiiLang` (ちい言語) はなんかちいさくてかわいいプログラミング言語です。

## インストール

```sh
go install github.com/aethiopicuschan/chiilang@latest
```

## 文法

- ワァ: ポインタをインクリメント
- ワワ: ポインタをデクリメント
- テメ: ポインタの指す値をインクリメント
- ヤンパパ: ポインタの指す値をデクリメント
- ンショ: ポインタの指す値を出力
- イヤッ: 入力をポインタの指す値に代入
- チャル: ポインタが指す値が0なら対応するジャンプ先に移動
- ジクジル: ポインタが指す値が0でないなら対応するジャンプ先に移動

なお、半角カナも受け付けます。

## 使い方

```sh
# 対話式 exitかquitで終了
$ chiilang

# ファイルから実行
$ chiilang run helloworld.chii
```

## おまけ

このリポジトリが提供するインタプリタは実際のところ純粋なBrainfuckインタプリタです。ただし、使用するコマンドを書き換える関数をインタプリタが持つことによって、上記の文法での実行が可能になっています。同様に `interpreter` パッケージを使うことで、あなた独自の言語を作成することもできます。詳しくは[Go Reference](https://pkg.go.dev/github.com/aethiopicuschan/chiilang)を参照してください。
