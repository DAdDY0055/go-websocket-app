package main

// clientはチャットを行なっている1人のユーザーを表す
type room struct {
	// forwordは他のクライアントに転送するためのメッセージを保持するチャンネル
	forword chan []byte
}
