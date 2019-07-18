package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func index(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl := template.Must(template.New("form").ParseFiles("add.html"))
		err := tmpl.Execute(w, nil)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}
func insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		tmpl := template.Must(template.New("result").ParseFiles("data.html"))

		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		nama := r.FormValue("nama")
		alamat := r.FormValue("alamat")
		jumlah := r.FormValue("jumlah")
		estimasi := r.FormValue("estimasi")

		data := map[string]string{"nama": nama, "alamat": alamat, "jumlah": jumlah, "estimasi": estimasi}

		if err := tmpl.Execute(w, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}
func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/process", insert)

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)

	db, err := sql.Open("postgres",
		"postgresql://root@conan:26257/mahasiswa_db?sslmode=disable")
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}
	defer db.Close()

}
