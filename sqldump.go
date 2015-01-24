/*
export GOPATH=~/bin/sqldump/
go get github.com/go-sql-driver/mysql
go run sqldump.go 
*/

package main
 
import ("fmt"
        "net/http"
 	"database/sql"
      _ "github.com/go-sql-driver/mysql"
  	"os"
 )

var base_url    = "http://localhost"
var user        = "go_user"
var pw 		= "mypassword"
var host        = "localhost"
var port        = "3306"
var database	= "information_schema"


func favicon(w http.ResponseWriter, r *http.Request) {
	http.StatusText(404)
}

func router(w http.ResponseWriter, r *http.Request) {
	
	if r.URL.Path == "/" {
		home(w,r)
	}else{
	      	parray := url2array(r)
                fmt.Fprintln(w, parray)
         
	        switch len(parray) {
		case 1:  dumpdb(w,r,parray)
		case 2:  dumptable(w,r,parray)
		case 3:  dumprecord(w,r,parray)}
	}
}


func home(w http.ResponseWriter, r *http.Request) {
         conn, err := sql.Open("mysql", dsn(user, pw, database))
         if err != nil {fmt.Println(err);os.Exit(1)}
 
         statement, err := conn.Prepare("show databases") 
         if err != nil {fmt.Println(err);os.Exit(1)}

         rows, err := statement.Query()
         if err != nil {fmt.Println(err);os.Exit(1)}         

         for rows.Next() {
                 var field string
                 rows.Scan(&field)
                 fmt.Fprintln(w, "DB :", field)
         }
         conn.Close()
}


func main() {
    http.HandleFunc("/favicon.ico", favicon)
    http.HandleFunc("/", router)
    fmt.Println("Listening at localhost:8080")
    http.ListenAndServe(":8080", nil)
}


