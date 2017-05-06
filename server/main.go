// Server of KVDB https://github.com/genesem/kvdb
// part of KVDB database. copyright(C) GeneSemerenko

package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

// Operators mapped: GET SET REMOVE KEYS

func handler(w http.ResponseWriter, r *http.Request) {
	//qpath := r.URL.Path[1:]
	qkey := r.URL.Query().Get("q")

	switch strings.ToUpper(r.Method) { // process http verbs:

	case "GET":
		var val string
		var ok bool
		subk := r.URL.Query().Get("k") // subkey

		if len(subk) == 0 {
			val, ok = hdGet(w, r, qkey)
		} else {
			val, ok = hdGetSubk(w, r, qkey, subk)
		}

		if !ok {
			w.WriteHeader(http.StatusNotFound) // 404 ERROR CODE
			fmt.Fprintf(w, "key %s not found", qkey)
		} else {
			fmt.Fprintf(w, "%s", val)
		}

	case "DELETE":

		if ok := hdDel(w, r, qkey); !ok {
			w.WriteHeader(http.StatusNotFound) // 404 if cant delete
			fmt.Fprintf(w, "Can`t delete key(s) %s", qkey)
		}
		//else {fmt.Fprintf(w, "DELETED key==%s", qkey) }

	case "POST": // add error check etc
		var ttl int // 0
		ttl, _ = strconv.Atoi(r.URL.Query().Get("t"))
		if ok := hdPost(w, r, qkey, ttl); !ok { // process post
			w.WriteHeader(http.StatusNotFound) // 404 cant add the key
			fmt.Fprintf(w, "Can`t add key %s with ttl==%d", qkey, ttl)
		}
		//else {w.WriteHeader(http.StatusCreated)} // 201 commented means: 200

	case "HEAD": // return list of keys matched the mask (regex)

		val, ok := hdHead(w, r, qkey)
		if !ok {
			w.WriteHeader(http.StatusNotFound) // 404 ERROR CODE
			fmt.Fprintf(w, "no keys found for the pattern: %s", qkey)
		} else {
			fmt.Fprintf(w, "%s", val)
		}

	default: //PUT verb are also here
		w.WriteHeader(http.StatusNotImplemented) // 501
		fmt.Fprintf(w, "unknown method: %s\n", r.Method)
		return
	}

	//fmt.Printf("qpath==[%s], qkey==[%s]\n", qpath, qkey) //html.EscapeString()
}

// entry point for server
func main() {

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "3000"
	}

	go func() { // watcher cleans database keys every 500 Milliseconds
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
