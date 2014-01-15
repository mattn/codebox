package main

import (
    "github.com/hoisie/web"
    "os"
)

func main() {
    web.Get("/", func() string {
        return "hello world"
    })
    web.Run(":" + os.Getenv("PORT"))
}
