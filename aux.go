package main
 
import ("strings"
        "net/http")

/*
func linkdown (
// will create a link one level deeper

*/

func dsn (user string, pw string, db string) string{
	return user + ":" + pw + "@tcp(" + host + ":" + port + ")/" + db 
}

func url2array(r *http.Request) []string {
    path := r.URL.Path
    path = strings.TrimSpace(path)
    if strings.HasPrefix(path, "/") {path = path[1:]}
    if strings.HasSuffix(path, "/") {path = path[:len(path) - 1]}
    return strings.Split(path, "/")
}

