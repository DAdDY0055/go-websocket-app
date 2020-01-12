package main

import (
	"log"
	"net/http"	
	"text/template"
	"path/filepath"
	"sync"
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
	t.templ.Execute(w, nil) // t.templ.Executeの戻り値はチェックするべき
}

func main() {
	r := newRoom()
	http.Handle("/", &templateHandler{filename: "chat.html"})
	http.Handle("/room", r)

	// チャットルームを開始
	go r.run()

	// Webサーバーを起動
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
