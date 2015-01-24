package main


/* TODO
 * correct columns
 * turn into a generic functions
 * prevent sql injection
 * provide links
 * 
 */

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"os"
)

//  Dump all tables of a database
func dumpdb(w http.ResponseWriter, r *http.Request, parray []string) {

	database := parray[0]
	conn, err := sql.Open("mysql", dsn(user, pw, database))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	statement, err := conn.Prepare("show tables")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	rows, err := statement.Query()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Fprintln(w, rows)
	var n = 1
	for rows.Next() {
		var field string
		rows.Scan(&field)
		fmt.Fprintln(w, n, field)
		n = n + 1
	}
	conn.Close()
}

//  Dump all records of a table, one per line
func dumptable(w http.ResponseWriter, r *http.Request, parray []string) {

	database := parray[0]
	table := parray[1]

	conn, err := sql.Open("mysql", dsn(user, pw, database))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	statement, err := conn.Prepare("select * from " + table)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	rows, err := statement.Query()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Fprintln(w, rows)

	var n = 1
	for rows.Next() {
		var field string
		rows.Scan(&field)
		fmt.Fprintln(w, n, field)
	}
	conn.Close()
}

// Shows selection of databases at top level 
func home(w http.ResponseWriter, r *http.Request) {
	conn, err := sql.Open("mysql", dsn(user, pw, database))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	statement, err := conn.Prepare("show databases")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	rows, err := statement.Query()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for rows.Next() {
		var field string
		rows.Scan(&field)
		fmt.Fprintln(w, "DB :", field)
	}
	conn.Close()
}



// Dump all fields of a record, one column per line
func dumprecord(w http.ResponseWriter, r *http.Request, parray []string) {

	// http://go-database-sql.org/varcols.html

	database := parray[0]
	table := parray[1]
	rec := parray[2]

	conn, err := sql.Open("mysql", dsn(user, pw, database))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	statement, err := conn.Prepare("select * from " + table + "limit " + rec + ",1")
	statement, err = conn.Prepare("show colums from " + table)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	rows, err := statement.Query()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Fprintln(w, rows)
	var n = 1
	for rows.Next() {
		var field string
		rows.Scan(&field)
		fmt.Fprintln(w, n, field)
	}
	conn.Close()
}
