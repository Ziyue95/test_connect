package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	// connect to a DB

	// create a connect pool conn
	conn, err := sql.Open("pgx", "host=localhost port=5433 dbname=test_connect user=home password=")
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to connect: %v\n", err))
	}
	defer conn.Close()

	log.Println("Connected to database!")

	// test my connect
	err = conn.Ping()
	if err != nil {
		log.Fatal("Can not ping database")
	}

	log.Println("Pinged database!")

	// get rows from table
	err = getAllRows(conn)
	if err != nil {
		log.Fatal(err)
	}

	// insert a row
	query := `insert into users2 (first_name, last_name) values($1, $2)`
	_, err = conn.Exec(query, "Tom", "Dwan")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Inserted a row.")

	// get rows from table again
	err = getAllRows(conn)
	if err != nil {
		log.Fatal(err)
	}

	// update a row
	stmt := `update users2 set first_name = $1 where id = $2`
	_, err = conn.Exec(stmt, "Jackie", 8)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Updated one or more rows.")

	// get rows from table again
	err = getAllRows(conn)
	if err != nil {
		log.Fatal(err)
	}

	// get one row by id
	query = `select id, first_name, last_name from users2 where id = $1`

	var firstName, lastName string
	var id int

	rows := conn.QueryRow(query, 5)
	err = rows.Scan(&id, &firstName, &lastName)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("QueryRow returns", id, firstName, lastName)

	// delete a row
	query = `delete from users2 where id = $1`
	_, err = conn.Exec(query, 7)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Deleted a row.")

	// get rows from table again
	err = getAllRows(conn)
	if err != nil {
		log.Fatal(err)
	}

}

func getAllRows(conn *sql.DB) error {
	rows, err := conn.Query("select id, first_name, last_name from users2")
	if err != nil {
		log.Println(err)
		return err
	}
	defer rows.Close() // important! Close rows after function return

	var firstName, lastName string
	var id int

	for rows.Next() {
		err := rows.Scan(&id, &firstName, &lastName)
		if err != nil {
			log.Println(err)
			return err
		}
		fmt.Println("Record is", id, firstName, lastName)
	}
	if err = rows.Err(); err != nil {
		log.Fatal("Error scanning rows", err)
	}

	fmt.Println("------------------------")

	return nil
}
