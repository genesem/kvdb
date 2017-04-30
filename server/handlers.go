// operations handlers
// implemented operations: GET SET REMOVE
// not implemented: KEYS

package main

import (
	"encoding/base64"
	"io/ioutil"
	"log"
	"net/http"
)

//i, ok := m["route"]
//In this statement, the first value (i) is assigned the value stored under the key "route".
//If that key doesn't exist, i is the value type's zero value (0). The second value (ok) is a bool that is true if the key exists in the map, and false if not.

// Get key
func hdGet(w http.ResponseWriter, req *http.Request, key string) (val string, ok bool) {

	ret, ok := DataBase[key]
	if ok {
		switch ret.Type { // set type prefix to ret result
		case TStr: //`string`
			val = "+"
		case TList: //`[]string` aks list
			val = "$"
		case TDict: //`map[string]string` aka dict
			val = "#"
		default: // error prefix
			val = "-"
		}
		val += base64.URLEncoding.EncodeToString(ret.Val)
		log.Printf("hdGet()2:\tret:%v\nkey:%s\tok==%t\n", ret.Val, key, ok)

		return val, true
	}
	return "", false

}

// Delete key
func hdDel(w http.ResponseWriter, req *http.Request, key string) bool {

	delete(DataBase, key)
	log.Printf("\nDelete key: %s\n", key)
	_, exist := DataBase[key] // check if key exists, if so its error
	return !exist
}

// Append key to database with ttl
func hdPost(w http.ResponseWriter, req *http.Request, key string, ttl int64) {

	defer req.Body.Close()
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	log.Printf("IN:hdPost append data: %v\n", data)
	DataBase[key] = &DBItem{ttl, TList, data}

}
