package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

type IndexData struct {
    Title string
    Content string
}

func test(c *gin.Context) {
    data:= new(IndexData)
    data.Title = "HomePage"
    data.Content = "My first homepage"
    c.HTML(http.StatusOK, "index.html", data)
}


func main() {
    server := gin.Default()
    server.LoadHTMLGlob("template/*")
    server.GET("/", test)
    server.Run(":8888")
}
