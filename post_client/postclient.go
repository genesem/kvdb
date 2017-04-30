package main

import (
	"bytes"
	cd "clientcodecs"
	"fmt"
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

	resp, err := http.Post(fmt.Sprintf(baseurl, "key1", 0), "application/octet-stream", bytes.NewBuffer(data))
	fmt.Printf("Resp==%v,  err==%v\n", resp.Status, err)
	resp.Body.Close()

	resp, err = http.Post(fmt.Sprintf(baseurl, "key2", 64), "application/octet-stream", bytes.NewBuffer(data))
	fmt.Printf("Resp==%v,  err==%v\n", resp.Status, err)
	resp.Body.Close()

	resp, err = http.Post(fmt.Sprintf(baseurl, "key3", 90), "application/octet-stream", bytes.NewBuffer(data))
	fmt.Printf("Resp==%v,  err==%v\n", resp.Status, err)
	resp.Body.Close()

	resp, err = http.Post(fmt.Sprintf(baseurl, "key4", 120), "application/octet-stream", bytes.NewBuffer(data))
	fmt.Printf("Resp==%v,  err==%v\n", resp.Status, err)
	resp.Body.Close()

	resp, err = http.Post(fmt.Sprintf(baseurl, "key5", 180), "application/octet-stream", bytes.NewBuffer(data))
	fmt.Printf("Resp==%v,  err==%v\n", resp.Status, err)
	resp.Body.Close()
}
