package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Wellcome to the home page.")
	fmt.Printf("Hit the home page endpoint")
}

func users() []*User {
	// Open up db connection
	db, err := sql.Open("mysql", "test_user:secret@tcp(db:3306)/test_database")

	// Check if error with db connection
	if err != nil {
		panic(err)
	}

	defer db.Close()

	// Excute query
	results, err := db.Query("SELECR * FROM users")

	if err != nil {
		panic(err)
	}

	var users []*User
	for results.Next() {
		var u User

		// Scan each result
		err = results.Scan(&u.ID, &u.Name)
		if err != nil {
			panic(err)
		}

		users = append(users, &u)
	}

	return users
}

func userPage(w http.ResponseWriter, r *http.Request) {
	users := users()

	fmt.Println("Hit the users oage endpoint")
	json.NewEncoder(w).Encode(users)
}

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/users", userPage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
