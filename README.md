# WebSocketを使ったチャットアプリケーション
by Go言語によるWebアプリケーション開発
https://www.oreilly.co.jp/books/9784873117522/

# USAGE

```
$ go build -o chat
```

```
$ ./chat -addr=":3000"
```
→ `3000`の部分は自由に選択可能



※1.3まで
```
$ go run main.go client.go room.go
```
→ 複数指定してrunしないとエラー
