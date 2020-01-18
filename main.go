package main

import (
	"log"
	"net/http"	
	"text/template"
	"path/filepath"
	"sync"
	"flag"
)

// templは1つのテンプレートを表す
type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

// ServeHTTPはHTTPリクエストを処理する
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = 
			template.Must(template.ParseFiles(filepath.Join("templates", 
				t.filename)))
	})
	t.templ.Execute(w, r) // t.templ.Executeの戻り値はチェックするべき
}

func main() {
	var addr = flag.String("addr", ":8080", "アプリケーションのアドレス")
	flag.Parse() // フラグを解釈する
	r := newRoom()
	http.Handle("/", &templateHandler{filename: "chat.html"})
	http.Handle("/room", r)

	// チャットルームを開始
	go r.run()

	// Webサーバーを起動
	log.Println("Webサーバーを開始します。ポート：", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
