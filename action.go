package main

/*
<form  action="/login">
   <label for="user">User</label><input type="text"     id="user" name="user"><br>
   <label for="pass">Pass</label><input type="password" id="pass" name="pass"><br>
   <label for="host">Host</label><input type="text"     id="host" name="host" value="localhost"><br>
   <label for="port">Port</label><input type="text"     id="port" name="port" value="3306"><br>
   <button type="submit">Query</button>
*/

import (
	"fmt"
	"net/http"
	//	"html/template"
)

func actionSelect(w http.ResponseWriter, r *http.Request, database string, table string) {
	fmt.Fprintln(w, "TODO select")
}

func actionInsert(w http.ResponseWriter, r *http.Request, database string, table string) {
	fmt.Fprintln(w, "TODO insert")
}
