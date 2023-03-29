package main

import (
    "os"
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    port := os.Getenv("PORT")

    if port == "" {
        port = "8000"
    }

    r := gin.Default()
    
    r.Static("/static", "resources/static")
    r.LoadHTMLGlob("resources/*.html")



    r.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.html", gin.H{
            "title": "Home Page",
            "storage_url": "https://ik.imagekit.io/ugivvozi5c",
            "static_path": "/static",
        })
    })

    r.Run(fmt.Sprintf(":%s", port))
}