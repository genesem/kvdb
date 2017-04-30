package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

// Operators mapped: GET SET REMOVE KEYS
func handler(w http.ResponseWriter, r *http.Request) {

	qpath, qkey := r.URL.Path[1:], r.URL.Query().Get("q")

	switch strings.ToUpper(r.Method) { // process http verbs:

	case "GET":
		v, ok := hdGet(w, r, qkey)
		if !ok {
			w.WriteHeader(http.StatusNotFound) // 404 ERROR CODE
			fmt.Fprintf(w, "key %s not found", qkey)
		} else {
			fmt.Fprintf(w, "%s", v)
		}

	case "DELETE":

		if ok := hdDel(w, r, qkey); !ok {
			w.WriteHeader(http.StatusNotFound) // 404 if cant delete
			fmt.Fprintf(w, "Can`t delete key %s", qkey)
		} else {
			fmt.Fprintf(w, "DELETED key==%s", qkey)
		}

	case "POST": // add error check etc
		ttl := time.Now().UnixNano()
		hdPost(w, r, qkey, ttl) // process post
		//w.WriteHeader(http.StatusCreated) // 201  or leave 200

	default: //head, put is also here
		w.WriteHeader(http.StatusNotImplemented) // 501
		fmt.Fprintf(w, "unknown method: %s\n", r.Method)
		return
	}

	fmt.Printf("qpath==[%s], qkey==[%s]\n", qpath, qkey) //html.EscapeString()
}

// entry point for the server
func main() {

	work()

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "3000"
	}

	go func() { // watcher cleans database keys every 250 Milliseconds
		for {
			select {
			case <-time.After(500 * time.Millisecond): // for debug == 500ms
				CleanDB()
			}
		}
	}()

	http.HandleFunc("/", handler)
	print("Server started at the port :", port, "\n\n")
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		println("ListenAndServe: ", err)
	}

}
