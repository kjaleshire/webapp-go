package main

import (
    "log"
    "net/http"

    "github.com/codegangsta/negroni"
)

func main() {
    n := negroni.New(
        negroni.NewRecovery(),
        negroni.HandlerFunc(MyMiddleware),
        negroni.NewLogger(),
        negroni.NewStatic(http.Dir("public")),
    )

    n.Run(":8080")
}

func MyMiddleware(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
    log.Println("Logging the start")

    if r.URL.Query().Get("password") == "supersecret" {
        next(rw, r)
    } else {
        http.Error(rw, "Not Authorized", 401)
    }

    log.Println("Logging the end")
}
