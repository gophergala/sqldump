package main


import ("fmt"
        "net/http"
 	"database/sql"
      _ "github.com/go-sql-driver/mysql"
  	"os"
 )




func dumpdb(w http.ResponseWriter, r *http.Request, parray []string) {

         database := parray[0]
         conn, err := sql.Open("mysql", dsn(user, pw, database))
         if err != nil {fmt.Println(err);os.Exit(1)}
 
         statement, err := conn.Prepare("show tables") 
         if err != nil {fmt.Println(err);os.Exit(1)}

         rows, err := statement.Query()
         if err != nil {fmt.Println(err);os.Exit(1)}         

	 fmt.Fprintln(w,rows)
	 var n = 1
         for rows.Next() {
                 var field string
                 rows.Scan(&field)
                 fmt.Fprintln(w, n, field)
                 n = n +1 
         }
         conn.Close()
}

func dumptable(w http.ResponseWriter, r *http.Request, parray []string) {

         database := parray[0]
         table	  := parray[1]

         conn, err := sql.Open("mysql", dsn(user, pw, database))
         if err != nil {fmt.Println(err);os.Exit(1)}
         if err != nil {fmt.Println(err);os.Exit(1)}
 
         statement, err := conn.Prepare("select * from " + table) 
         if err != nil {fmt.Println(err);os.Exit(1)}

         rows, err := statement.Query()
         if err != nil {fmt.Println(err);os.Exit(1)}         

	 fmt.Fprintln(w,rows)

	 var n = 1
         for rows.Next() {
                 var field string
                 rows.Scan(&field)
                 fmt.Fprintln(w, n, field)
         }
         conn.Close()
}

func dumprecord(w http.ResponseWriter, r *http.Request, parray []string) {
 
         database := parray[0]
         table	  := parray[1]
         rec	  := parray[2]

         conn, err := sql.Open("mysql", dsn(user, pw, database))
         if err != nil {fmt.Println(err);os.Exit(1)}
 
         statement, err := conn.Prepare("select * from " + table + "limit " + rec + ",1") 
         statement, err = conn.Prepare("show colums from " + table) 
         if err != nil {fmt.Println(err);os.Exit(1)}

         rows, err := statement.Query()
         if err != nil {fmt.Println(err);os.Exit(1)}         

	 fmt.Fprintln(w,rows)
	 var n = 1
         for rows.Next() {
                 var field string
                 rows.Scan(&field)
                 fmt.Fprintln(w, n, field)
         }
         conn.Close()
}

