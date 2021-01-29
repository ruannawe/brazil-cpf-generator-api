package main

import (
    "github.com/kataras/iris/v12"
    "github.com/mvrilo/go-cpf"
)

func main() {
    app := iris.New()

    booksAPI := app.Party("/")
    {
        booksAPI.Use(iris.Compression)
        booksAPI.Get("/", GenerateCNPJ)
    }

    app.Listen(":8080")
}

func GenerateCNPJ(ctx iris.Context) {
    ctx.JSON(cpf.GeneratePretty())
}
