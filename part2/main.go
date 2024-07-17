package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"todoAPp/part2/util"
)

func main() {
	db, err := sql.Open("postgres", util.CONNECT_STR)
	if err != nil {
		log.Fatal(err)
	}
	//rows, err := db.Query("SELECT * FROM todo_list")
	defer db.Close()
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	log.Println("connecting successfully....")
	var todothing util.ToDOThing
	var todoArray []util.ToDOThing

	rows, err := db.Query("SELECT id, name, period, status FROM todo_list")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&todothing.Id, &todothing.Name, &todothing.Period, &todothing.Status)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(todothing)
		todoArray = append(todoArray, todothing)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	log.Println(todoArray)
}
