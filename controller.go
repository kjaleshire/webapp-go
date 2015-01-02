package main

import(
    "net/http"

    "gopkg.in/unrolled/render.v1"
)

type Action func(rw http.ResponseWriter, r *http.Request) error

type AppController struct{}

type MyController struct {
    AppController
    *render.Render
}

func (c *AppController) Action(a Action) http.Handler {
    return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
        if err := a(rw, r); err != nil {
            http.Error(rw, err.Error(), 500)
        }
    })
}

func (c *MyController) Index(rw http.ResponseWriter, r *http.Request) error {
    c.JSON(rw, 200, map[string]string{"Hello":"JSON"})
    return nil
}

func main() {
    c := &MyController{Render: render.New(render.Options{})}
    http.ListenAndServe(":8080", c.Action(c.Index))
}
