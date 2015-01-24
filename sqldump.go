/* TODO
 * login and sessions
 */


package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

var base_url = "http://localhost"
var user = "go_user"
var pw = "mypassword"
var host = "localhost"
var port = "3306"
var database = "information_schema"

func favicon(w http.ResponseWriter, r *http.Request) {
	http.StatusText(404)
}

func router(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path == "/" {
		home(w, r)
	} else {
		parray := url2array(r)

//		fmt.Fprintln(w, parray)

		switch len(parray) {
		case 1:
			dumpdb(w, r, parray)
		case 2:
			dumptable(w, r, parray)
		case 3:
			dumprecord(w, r, parray)
		}
	}
}


func main() {
	http.HandleFunc("/favicon.ico", favicon)
	http.HandleFunc("/", router)
	fmt.Println("Listening at localhost:8080")
	http.ListenAndServe(":8080", nil)
}
