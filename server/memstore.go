package main

import (
	"fmt"
	"time"
)

type ItemType int

const ( // possible types:
	TStr  ItemType = iota //`string` тип строка
	TList                 //`[]string` тип список. значение может быть получено по index
	TDict                 //`map[string]string` тип словарь значение может быть получено по ключу
)

type DBItem struct {
	TTL  int64 // this is time in seconds when key is expired i.e. time.Now().Unix()+seconds
	Type ItemType
	Val  []byte
}

// Returns true if the item has expired
func (it DBItem) Expired() bool {
	if it.TTL == 0 {
		return false
	}
	return time.Now().Unix() > it.TTL
}

// Set TTL of the item by seconds
func (it DBItem) SetTTL(sec int) { // if sec == 0 set ttl = 0 (means forver)
	if sec == 0 {
		it.TTL = int64(0)
		return
	}
	it.TTL = time.Now().Unix() + int64(sec)
}

var DataBase = make(map[string]*DBItem, 1<<16) // Allocate memory to 65536 items

// Cleans some random keys
func CleanDB() {

	//print("\rClean DB..")

	// Loop over map and append keys to empty slice.
	keys := []string{}
	var i int
	for key, _ := range DataBase { // value skipped, order is random

		if DataBase[key].Expired() {
			delete(DataBase, key)
			keys = append(keys, key) // here can be 'delete'
		}

		if i > 3 { // 3 gives 5 keys [0..4] // must be 20-25
			break // stop the loop
		}
		i++
	}

	if len(keys) > 0 {
		fmt.Printf("Keys deleted: %v\n", keys)
	}

}
