package main

import (
    "fmt"
    "html"
    "math/big"
    "net/http"
)

type Server struct{}

func main() {
    http.ListenAndServe("localhost:8000", Server{})
}
func (Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    left := r.FormValue("left")
    right := r.FormValue("right")
    op := r.FormValue("op")

    leftInt := &big.Int{}
    rightInt := &big.Int{}

    _, leftres := leftInt.SetString(left, 10)
    _, rightres := rightInt.SetString(right, 10)
///////////////////////////////////////////////
    var result string
    if leftres && rightres {
        resultInt := &big.Int{}
        switch op {
        case "add":
            resultInt.Add(leftInt, rightInt)
        case "sub":
            resultInt.Sub(leftInt, rightInt)
        case "multi":
            resultInt.Mul(leftInt, rightInt)
        case "div":
            resultInt.Div(leftInt, rightInt) 
        }
        result = resultInt.String()
    }
    switch left{
    case "":
        result ="左項目何か入れろ"
    }
    switch right{
    case "":
        result ="右項目何か入れろ"
    }
    h := `
            <html>
                <head>
                    <title>電卓アプリ</title>
                </head>
                <body>
                    <form action ="/" method="post">
                        左項目：<input type="text" name="left"><br>
                        右項目：<input type="text" name="right"><br>
                        演算子：
                        <input type="radio" name="op" value="add" > +
                        <input type="radio" name="op" value="sub" > -
                        <input type="radio" name="op" value="multi" > ×
                        <input type="radio" name="op" value="div" > ÷
                        <br><input type="submit" name="送信"><hr>

                        [フォームの入力値]<br>
                        左項目：` + html.EscapeString(left) + `<br>
                        右項目：` + html.EscapeString(right) + `<br>
                        演算子：` + html.EscapeString(op) + `<br>
                        演算結果：` + html.EscapeString(result) + `<br>
                    </form>
                </body>
            </html>   `
    fmt.Fprint(w, h)
}
