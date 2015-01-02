package main

import(
    "html/template"
    "net/http"
    "path"
)

type Book struct {
    Title   string
    Author  string
}

func main() {
    http.HandleFunc("/", ShowBooks)
    http.ListenAndServe(":8080", nil)
}

func ShowBooks(w http.ResponseWriter, r *http.Request) {
    book := Book{"Building Web Apps with Go", "Jeremy Saenz"}

    file_path := path.Join("templates", "template.html")
    tmpl, error := template.ParseFiles(file_path)
    if error != nil {
        http.Error(w, error.Error(), http.StatusInternalServerError)
        return
    }

    if error := tmpl.Execute(w, book); error != nil {
        http.Error(w, error.Error(), http.StatusInternalServerError)
    }

}
