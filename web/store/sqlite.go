package store

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type SqliteStore interface {
}

type sqllitestore struct {
}

func CreateAndPrint() {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatal(err)
		return
	}

	table := `create table userinfo(
		id integer primary key autoincrement,
		username varchar(50) not null,
		password varchar(128) not null,
		create_time date not null
	);
	`
	_, dberr := db.Exec(table)
	if dberr != nil {
		log.Fatal(dberr)
		return
	}

	stmt, perr := db.Prepare("insert into userinfo (username, password,create_time)values(?,?,?)")
	if perr != nil {
		log.Fatal(perr)
		return
	}
	res, err := stmt.Exec("demo", "123456", "now()")
	if err != nil {
		log.Fatal(perr)
		return
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal(perr)
		return
	}
	fmt.Printf("Id: %v", id)

	rows, err := db.Query("SELECT * FROM userinfo")
	if err != nil {
		log.Fatal(perr)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var username string
		var password string
		var created time.Time
		err = rows.Scan(&id, &username, &password, &created)
		if err != nil {
			panic(err)
		}

		fmt.Print(id)
		fmt.Print(username)
		fmt.Print(password)
		fmt.Print(created)
		fmt.Println()
	}

	db.Close()
}
