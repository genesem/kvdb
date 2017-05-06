package main

import (
	cd "clientcodecs"
	"fmt"
	//cd "github.com/genesem/kvdb/clientcodecs"
	"io/ioutil"
	"net/http"
	"time"
)

const baseurl = "http://localhost:3000/?q=%s"

func main() {
	strkeys := []string{"key1", "key2", "key3", "key4", "key5", "key6", "key7"}
	done := make(chan string)

	for ic := 0; ic < 1; ic++ { // loop to stress: 1000
		start := time.Now()

		for _, u := range strkeys {
			go func(u string) {

				resp, err := http.Get(fmt.Sprintf(baseurl, u))
				defer resp.Body.Close() // close resp when finished

				if err != nil {
					done <- u + " " + err.Error()
				} else {

					objx := []string{}

					if resp.StatusCode == http.StatusOK {
						body, _ := ioutil.ReadAll(resp.Body) // process err later
						fmt.Printf("got body:`%s`\n", body)

						switch body[0] { // set type prefix to ret result
						case '+': //`string`
							fmt.Printf("got string\n")
						case '$': //`[]string` aks list
							fmt.Printf("got list\n")
						case '#': //`map[string]string` aka dict
							fmt.Printf("got dict\n")
						case '-': // error
							fmt.Printf("got error\n")
						default: // return format error
						}

						cd.Decode(&objx, body[1:]) // skip 1st char as its type sign

					} else {
						fmt.Printf("key %s not found\n", u)
					}
					fmt.Printf("result for key(%s):`%v`\n", u, objx)

					done <- u + " " + resp.Status

				}
			}(u)
		}

		for _ = range strkeys {
			fmt.Println(<-done, time.Since(start))
		}

	} // end loop

}
