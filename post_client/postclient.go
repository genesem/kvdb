package main

import (
	"bytes"
	//cd "clientcodecs"
	"fmt"
	cd "github.com/genesem/kvdb/clientcodecs"
	"net/http"
)

const (
	USD int = iota
	EUR
	GBP
	RUR
)

const baseurl = "http://localhost:3000/?q=%s&t=%d"

func main() {

	// data to store
	list1 := []string{USD: "US", EUR: "EUR", GBP: "GBP", RUR: "RUB"} // list
	data := cd.Encode(list1)

	/*	client := &http.Client{} // future use
		r, _ := http.NewRequest("POST", fmt.Sprintf(baseurl, key, ttl), bytes.NewBufferString(data))
		//r.Header.Add("Authorization", "auth_token=\"XXXXXXX\"")
		r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		r.Header.Add("Content-Length", strconv.Itoa(len(data)))
		resp, _ := client.Do(r)
	*/

	strkeys := [...]string{"key1", "key2", "key3", "key4", "key5", "key6", "key7"}
	ttls := [...]int{0, 30, 60, 90, 120, 130, 180}

	for i, u := range strkeys {

		resp, err := http.Post(fmt.Sprintf(baseurl, u, ttls[i]), "application/octet-stream", bytes.NewBuffer(data))
		fmt.Printf("Resp==%v,  err==%v\n", resp.Status, err)
		resp.Body.Close()
	}

}
