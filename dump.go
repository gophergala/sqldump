package main

/* TODO
 * turn into more generic functions for printing into tables
 */

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"net/http"
	"strconv"
)



// Shows selection of databases at top level
func dumpHome(w http.ResponseWriter, r *http.Request) {

	user, pw, h, p := getCredentials(r)
	conn, err := sql.Open("mysql", dsn(user, pw, h, p, database))
	checkY(err)
	defer conn.Close()

	statement, err := conn.Prepare("show databases")
	checkY(err)

	rows, err := statement.Query()
	checkY(err)
	defer rows.Close()

	var n int = 1
	for rows.Next() {
		var field string
		rows.Scan(&field)
		tableDuo(w, href(r.URL.Host + "?" + "db=" + field, "["+strconv.Itoa(n)+"]"), field)
		n = n + 1
	}
}

//  Dump all tables of a database
func dumpTables(w http.ResponseWriter, r *http.Request, database string) {

	user, pw, h, p := getCredentials(r)
	conn, err := sql.Open("mysql", dsn(user, pw, h, p, database))
	checkY(err)
	defer conn.Close()

	statement, err := conn.Prepare("show tables")
	checkY(err)

	rows, err := statement.Query()
	checkY(err)
	defer rows.Close()

	var n int = 1
	for rows.Next() {
		var field string
		rows.Scan(&field)
		tableDuo(w, href(r.URL.Host + "?" + r.URL.RawQuery + "&t=" + field, "["+strconv.Itoa(n)+"]"), field)
		n = n + 1
	}
}

//  Dump all records of a table, one per line
func dumpRecords(w http.ResponseWriter, r *http.Request, database string, table string) {

	user, pw, h, p := getCredentials(r)

	conn, err := sql.Open("mysql", dsn(user, pw, h, p, database))
	checkY(err)
	defer conn.Close()

	statement, err := conn.Prepare("select * from " + template.HTMLEscapeString(table))
	checkY(err)

	rows, err := statement.Query()
	checkY(err)
	defer rows.Close()

	cols, err := rows.Columns()
	checkY(err)

	{ 
		// table head
		fmt.Fprint(w, lineA)
		tableHead(w, "#")
		for _, col := range cols {
			tableHead(w, col)
		}
		fmt.Fprint(w, lineO)
	}

	/*  credits:
	 * 	http://stackoverflow.com/questions/19991541/dumping-mysql-tables-to-json-with-golang
	 * 	http://go-database-sql.org/varcols.html
	 */

	raw := make([]interface{}, len(cols))
	val := make([]interface{}, len(cols))

	for i := range val {
		raw[i] = &val[i]
	}

	var n int = 1
	for rows.Next() {

		fmt.Fprint(w, lineA)
		tableCell(w, href(r.URL.Host + "?" + r.URL.RawQuery + "&n=" + strconv.Itoa(n), strconv.Itoa(n)))

		err = rows.Scan(raw...)
		checkY(err)

		for _, col := range val {
			if col != nil {
				tableCell(w, string(col.([]byte)))
			}
		}
		fmt.Fprint(w, lineO)
		n = n + 1
	}
}

// Dump all fields of a record, one column per line
func dumpFields(w http.ResponseWriter, r *http.Request, database string, table string, num string) {

	rec, err := strconv.Atoi(num)
	checkY(err)

	user, pw, h, p := getCredentials(r)
	conn, err := sql.Open("mysql", dsn(user, pw, h, p, database))
	checkY(err)
	defer conn.Close()

	statement, err := conn.Prepare("select * from " + template.HTMLEscapeString(table))
	checkY(err)

	rows, err := statement.Query()
	checkY(err)
	defer rows.Close()

	columns, err := rows.Columns()
	checkY(err)

	raw := make([]interface{}, len(columns))
	val := make([]interface{}, len(columns))

	for i := range val {
		raw[i] = &val[i]
	}

	var n int = 1

rowLoop:
	for rows.Next() {

		// unfortunately we have to iterate up to row of interest
		if n == rec {
			err = rows.Scan(raw...)
			checkY(err)

			for i, col := range val {
				if col != nil {
					tableDuo(w, columns[i], string(col.([]byte)))
				}
			}
			break rowLoop
		}
		n = n + 1
	}
}
