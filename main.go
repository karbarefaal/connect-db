package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "user=ali password=123456 dbname=test sslmode=disable port=5454")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users(
		id SERIAL PRIMARY KEY,
		username VARCHAR(50) NOT NULL UNIQUE,
		email VARCHAR(100) NOT NULL UNIQUE,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`)
	if err != nil {
		log.Fatal(err)
	}

	res, err := db.Exec(`INSERT INTO users (username,email) VALUES($1,$2 )`, "ali10", "sajjadia99722232@gmail.com")
	if err != nil {
		log.Fatal(err)
	}
	println(res.RowsAffected())

	//=================================================================================================================

	// res, err := db.Exec(`UPDATE users SET email=$1 WHERE username= $2`, "xxx@aaa", "ali")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(res)

	//=================================================================================================================

	// 	res, err := db.Exec(`DELETE FROM users WHERE username=$1 `, "ali")
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Println(res)

	//=================================================================================================================

	rows, err := db.Query(`SELECT id,username,email,created_at FROM users;`)

	defer rows.Close()

	for rows.Next() {
		var id int
		var username, email string
		var createdAt string
		err = rows.Scan(&id, &username, &email, &createdAt)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("id= %d, username = %s, email= %s, createdAt= %s\n", id, username, email, createdAt)

	}

}
