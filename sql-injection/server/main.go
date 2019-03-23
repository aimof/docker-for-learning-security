package main

import (
	"database/sql"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	user         = "user"
	rootPassword = "passwd"
	host         = "mysql"
	port         = "3306"
	dbName       = "injection"
	tableName    = "comment"
	db           = initDB()
)

func initDB() *sql.DB {
	time.Sleep(13 * time.Second)
	db, err := sql.Open("mysql", user+":"+rootPassword+"@tcp("+host+":"+port+")/"+dbName)
	if err != nil {
		log.Fatalln(err)
	}
	return db
}

func main() {
	defer db.Close()
	_, err := db.Query("CREATE TABLE IF NOT EXISTS " + dbName + "." + tableName + " (id INT, comment VARCHAR(1024))")
	if err != nil {
		db.Close()
		log.Fatalln(err)
	}
	_, err = db.Query("INSERT INTO " + dbName + "." + tableName + " VALUES (1, 'one'), (2, 'two'), (3, 'three'), (4, 'four'), (1024, 'forbidden'), (16384, 'hidden')")
	if err != nil {
		log.Fatalln(err)
	}
	http.HandleFunc("/", inputHandler)
	http.HandleFunc("/show", showHandler)
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func inputHandler(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("input.html")
	if err != nil {
		w.WriteHeader(500)
		return
	}
	defer file.Close()
	html, err := ioutil.ReadAll(file)
	if err != nil {
		w.WriteHeader(500)
	}
	w.Write(html)
}

func showHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./show.html")
	if err != nil {
		log.Println("cannot load template")
		w.WriteHeader(500)
	}
	type Item struct {
		Number int
		Value  string
	}
	num, ok := r.URL.Query()["number"]
	if !ok || len(num) == 0 {
		log.Println("Invalid URL")
		w.WriteHeader(500)
	}
	query := "SELECT * FROM " + tableName + " WHERE id <= " + num[0] + ";"
	rows, err := db.Query(query)
	if err != nil {
		log.Println(query)
		w.WriteHeader(500)
		return
	}
	defer rows.Close()
	Items := make([]Item, 0, 10)
	for rows.Next() {
		var num int
		var comment string
		if err := rows.Scan(&num, &comment); err != nil {
			log.Println(err)
			w.WriteHeader(500)
			return
		}
		Items = append(Items, Item{num, comment})
	}
	log.Println(Items)
	err = tmpl.Execute(w, Items)
	if err != nil {
		log.Println(Items)
		w.WriteHeader(500)
		return
	}
}
