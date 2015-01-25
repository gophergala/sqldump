package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

var base_url = "http://localhost"
var database = "information_schema"

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.StatusText(404)
}

func loginPageHandler(w http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(w, loginPage)
}

func pathHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path == "/" {
		fmt.Fprintln(w, "</p>")
		home(w, r)
	} else {
		parray := url2array(r)

		switch len(parray) {
		case 1:
			fmt.Fprintln(w, parray[0], "</p>")
			dumpdb(w, r, parray)
		case 2:
			fmt.Fprintln(w, parray[0], "/", parray[1],"</p>")
			dumptable(w, r, parray)
		case 3:
			fmt.Fprintln(w, parray[0], "/", parray[1], "/", parray[2],"</p>")
			dumprecord(w, r, parray)
		}
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	user , _ , host , port := getCredentials(r)

	if user != "" {
		fmt.Fprint(w, "<p>", linkDeeper(r.URL.Scheme + r.URL.Host, "exit", "[x]")," &nbsp; ", user + "@" + host + ":" + port, " &nbsp; ")
		// TODO remove this ugly hack
		pathHandler(w, r)
	} else {
		loginPageHandler(w, r)
	}
}

func main() {

	http.HandleFunc("/favicon.ico", faviconHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/exit", logoutHandler)
	http.HandleFunc("/", indexHandler)

	fmt.Println("Listening at localhost:8080")
	http.ListenAndServe(":8080", nil)
}
