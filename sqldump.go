package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	//	"net/url"
)

var base_url = "http://localhost"
var database = "information_schema"

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.StatusText(404)
}

func loginPageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, loginPage)
}

func helpHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// TODO help
	fmt.Fprintf(w, "Best viewed with cli-browser >= 6.0")
}

func dumpPath(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path == "/" {
		fmt.Fprintln(w, "</p>")
		fmt.Fprint(w, tableA)
		home(w, r)
		fmt.Fprint(w, tableO)
	} else {
		parray := url2array(r)

		fmt.Fprintln(w, tableA)
		switch len(parray) {
		case 1:
			fmt.Fprintln(w, parray[0], "</p>")
			dumpTables(w, r, parray)
		case 2:
			fmt.Fprintln(w, parray[0]+"."+parray[1], "</p>")
			dumpRecords(w, r, parray)
		case 3:
			fmt.Fprintln(w, parray[0]+"."+parray[1], "["+parray[2]+"]</p>")
			dumpFields(w, r, parray)
		}
		fmt.Fprintln(w, tableO)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	user, _, host, port := getCredentials(r)

	if user != "" {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		// TODO remove this ugly hack starting a line here and ending it somewhere else
		fmt.Fprint(w, href(r.URL.Scheme+r.URL.Host, "logout", "[X]"))
		fmt.Fprint(w, " &nbsp; ")
		fmt.Fprint(w, href(r.URL.Scheme+r.URL.Host, "help", "[?]"))
		fmt.Fprint(w, " &nbsp; ")
		fmt.Fprint(w, href(r.URL.Scheme+r.URL.Host, "", "[/]"))
		fmt.Fprint(w, " &nbsp; ")
		fmt.Fprint(w, user+"@"+host+":"+port)
		fmt.Fprint(w, " &nbsp; ")
		dumpPath(w, r)
	} else {
		v := r.URL.Query()
		Quser := v.Get("user")
		Qpass := v.Get("pass")
		Qhost := v.Get("host")
		Qport := v.Get("port")
		if Quser != "" && Qpass != "" {
			if Qhost == "" {
				Qhost = "localhost"
			}
			if Qport == "" {
				Qport = "3306"
			}
			setCredentials(w, Quser, Qpass, Qhost, Qport)
			http.Redirect(w, r, r.URL.Host, 302)
		} else {
			loginPageHandler(w, r)
		}
	}
}

func main() {

	http.HandleFunc("/favicon.ico", faviconHandler)
	http.HandleFunc("/help", helpHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/logout", logoutHandler)
	http.HandleFunc("/", indexHandler)

	fmt.Println("Listening at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
