// operations handlers
// implemented operations: GET SET REMOVE
// not implemented: KEYS

package main

import (
	"encoding/base64"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// hdHead: return list of keys matched the pattern
func hdHead(w http.ResponseWriter, req *http.Request, key string) (val string, ok bool) {

	r, _ := regexp.Compile(key) // ie 'ke??' or 'k*'
	var ret []string
	for k, _ := range DataBase { // iterate over db
		if r.MatchString(k) {
			ret = append(ret, k)
		}
	}
	//log.Println("result:", val)
	if len(ret) > 0 {
		return strings.Join(ret, ","), true
	}
	return "", false
}

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

	//log.Printf("\nDelete got string keys: %s\n", key)
	res := strings.Split(key, ",") // split keys by ,
	for i := range res {
		k := res[i] // better
		delete(DataBase, k)
		log.Printf("\nDelete key: %s\n", k)
		if _, ok := DataBase[k]; ok == true { // if key exists its an error!
			//log.Printf("\nDelete key TEST: %s, db: %v\n", k, DataBase[k])
			return false
		}
	}

	return true // all fine
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
