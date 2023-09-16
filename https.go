package main

import (
    // "fmt"
    // "io"
    "net/http"
    "log"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
    w.Header().Set("Content-Type", "text/plain")
    w.Write([]byte("This is an example https server(v1).\n"))
    // fmt.Fprintf(w, "This is an example server.\n")
    // io.WriteString(w, "This is an example server.\n")
}

func Home(w http.ResponseWriter, req *http.Request) {
    w.Header().Set("Content-Type", "text/plain")
    w.Write([]byte("Https test.Home Page(v1).\n"))
    // fmt.Fprintf(w, "This is an example server.\n")
    // io.WriteString(w, "This is an example server.\n")
}
func main() {
    http.HandleFunc("/hello", HelloServer)
    http.HandleFunc("/",Home)
    err := http.ListenAndServeTLS(":443", "tls.crt", "tls.key", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
