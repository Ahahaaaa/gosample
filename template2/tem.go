package main
 
import (
    "fmt"
    "html/template"
    "net/http"
)
 
func gtplHandler(w http.ResponseWriter, r *http.Request) {
 
    // テンプレートをパースする
    tpl := template.Must(template.ParseFiles("templates/sample.gtpl"))
 
    // テンプレートに出力する値をマップにセット
    values := map[string]string{
        "account": "user-1",
        "name":    "さとうこうや",
        "passwd":  "sample-pass",
    }
 
    // マップを展開してテンプレートを出力する
    if err := tpl.ExecuteTemplate(w, "sample.gtpl", values); err != nil {
        fmt.Println(err)
    }
}
 
func main() {
    // ルートへのリクエストを"gtplHandler"関数で処理する
    http.HandleFunc("/", gtplHandler)
 
    // localhost:8080でサーバー処理開始
    http.ListenAndServe(":8080", nil)
}