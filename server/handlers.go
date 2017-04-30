// operations handlers
// implemented operations: GET SET REMOVE
// not implemented: KEYS

package main

import (
	"encoding/base64"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

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
		//log.Printf("hdGet()2:\tret:%v\nkey:%s\tok==%t\n", ret.Val, key, ok)
		return val, true
	}
	return "", false

}

// get subkey if possible..
func hdGetSubk(w http.ResponseWriter, req *http.Request, key, subk string) (val string, ok bool) {

	ret, ok := DataBase[key]

	if ok {
		xdata := ret.Val

		switch ret.Type { // set type prefix to ret result
		case TStr: //`string`
			val = "+"
		case TList: //`[]string` aks list
			val = "$"
			ix, errx := strconv.Atoi(subk)
			if subk != "" && errx == nil { // need err processing and debug
				list1 := []string{}
				decode(&list1, ret.Val)
				xdata = encode(list1[ix])
			}
		case TDict: //`map[string]string` aka dict
			val = "#"
			if subk != "" { // need err processing and debug
				dict1 := map[string]string{}
				decode(&dict1, ret.Val)
				xdata = encode(dict1[subk])
			}

		default: // error prefix
			val = "-"
		}
		val += base64.URLEncoding.EncodeToString(xdata)
		log.Printf("hdGet()2:\tret:%v\nkey:%s\tok==%t\n", xdata, key, ok)

		return val, true
	}
	return "", false

}

// Delete key
func hdDel(w http.ResponseWriter, req *http.Request, key string) bool {

	delete(DataBase, key)
	log.Printf("\nDelete key: %s\n", key)
	_, exist := DataBase[key] // if key exists its an error!
	return !exist
}

// Append key to database with ttl
// ttl here is the number of seconds to live since NOW..
// ie. ttl=128 means two seconds of life to the key

func hdPost(w http.ResponseWriter, req *http.Request, key string, ttl int) bool {

	defer req.Body.Close()
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	_, err = base64.URLEncoding.Decode(data, data)
	if err == nil {
		//log.Printf("IN:hdPost append data: %v\n", data)
		var ttd int64 // def == 0
		if ttl != 0 {
			ttd = time.Now().Unix() + int64(ttl)
		}
		log.Printf("IN:hdPost ADD ttl==%d, ttd==%d", ttl, ttd)

		DataBase[key] = &DBItem{ttd, TList, data}
		return true
	}
	return false
}
