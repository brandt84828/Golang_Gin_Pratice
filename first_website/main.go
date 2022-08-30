package main

import (
    "log"
    "html/template"
    "net/http"
)

type IndexData struct {
    Title string
    Content string
}

func test(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`my first website`))
}

func test2(w http.ResponseWriter, r *http.Request) {
    str := `<!DOCTYPE html>
<html>
<head><title>首頁</title></head>
<body><h1>首頁</h1><p>我的第一個首頁</p></body>
</html>
`
    w.Write([]byte(str))
}

func test3(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("./index.html"))
    data := new(IndexData)
    data.Title = "HomePage"
    data.Content = "My first homepage"
    tmpl.Execute(w, data)
}


/*
HandleFunc 讓 server 知道當進來的 traffic 的 routing 為 / 時要執行 test 方法

ListenAndServe 有兩個參數分別為 address 與 handler:
address 為存取的 url 與 port，因為沒有指定 url 所以沒有填寫，只單純寫 port
本範例中沒有實作 handler，因此將它填寫為 nil，
*/


func main() {
    http.HandleFunc("/", test3)
    http.HandleFunc("/index", test3)
    err := http.ListenAndServe(":8888", nil)
    if err != nil {
	log.Fatal("ListenAndServe: ", err)
    }
}
