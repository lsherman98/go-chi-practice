package main

import (
    "fmt"
    "context"
    "github.com/lsherman98/go-chi-practice/application"
)

func main(){
    app := application.New()

    err := app.Start(context.TODO())
    if err != nil {
        fmt.Println("failed to start app:", err)
    }
}


