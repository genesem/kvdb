package main

import (
	"fmt"
	"time"
)

const (
	USD int = iota
	EUR
	GBP
	RUR
)

func work() {

	str1 := "#1: Hello this is a string one..." // string 1
	str2 := "#2: Hello this is a string two..." // string 2

	list1 := []string{USD: "US", EUR: "EUR", GBP: "GBP", RUR: "RUB"}    // list
	list2 := []string{"hello1", "hello2", "hello3", "hello4", "hello5"} // list2

	dict1 := map[string]string{"key1": "value1", "key2": "value2", "key3": "value3", "key4": "value4", "key5": "value5"}

	fmt.Printf("Original data str1:\t%#v\n", str1)
	fmt.Printf("Original data list1:\t%#v\n", list1)
	fmt.Printf("Original data dict1:\t%#v\n", dict1)

	// basic tests
	result := encode(list1)
	fmt.Printf("\nEncoding Result: %v\n", result)
	fmt.Printf("\nEncoding Result dict1: %#v\n", encode(dict1))

	DataBase["key1"] = &DBItem{time.Now().Unix() + 128, TList, result} // stub
	DataBase["key2"] = &DBItem{time.Now().Unix() + 128, TList, result} // stub
	DataBase["key3"] = &DBItem{time.Now().Unix() + 64, TList, result}  // stub
	DataBase["key4"] = &DBItem{time.Now().Unix() + 128, TList, result} // stub
	DataBase["key5"] = &DBItem{time.Now().Unix() + 256, TList, result} // stub

	DataBase["key6"] = &DBItem{0, TList, result}   // stub forever key
	DataBase["key7"] = &DBItem{0, TList, result}   // stub
	DataBase["key7"].TTL = time.Now().Unix() + 200 // 200 sec

	fmt.Printf("key1 ttl=%d\n", DataBase["key1"].TTL)
	fmt.Printf("key5 ttl=%d\n", DataBase["key5"].TTL)
	fmt.Printf("key6 ttl=%d\n", DataBase["key6"].TTL)
	fmt.Printf("key7 ttl=%d\n", DataBase["key7"].TTL)

	// store data to DB
	DataBase["list1"] = &DBItem{time.Now().Unix() + 128, TList, encode(list1)}
	DataBase["list2"] = &DBItem{time.Now().Unix() + 128, TList, encode(list2)}

	DataBase["str1"] = &DBItem{time.Now().Unix() + 128, TStr, encode(str1)}
	DataBase["str2"] = &DBItem{time.Now().Unix() + 128, TStr, encode(str2)}
	DataBase["dict1"] = &DBItem{time.Now().Unix() + 128, TDict, encode(dict1)}
	DataBase["dict2"] = &DBItem{time.Now().Unix() + 128, TDict, encode(dict1)}

	// show db state
	fmt.Printf("\nDatabase:\n%v\n", DataBase)

	delete(DataBase, "key2")
	delete(DataBase, "key4")
	delete(DataBase, "str2")
	delete(DataBase, "dict2")

	fmt.Printf("\nDatabase after delete keys:\n%v\n", DataBase)

	// test read

	// strings
	var str1x, str2x string
	decode(&str1x, DataBase["str1"].Val)
	decode(&str2x, DataBase["str1"].Val) // deleted
	fmt.Printf("Decoded Result str1: %#v\n", str1x)

	// lists
	var obj1x []string
	obj2x := []string{}

	decode(&obj1x, DataBase["list1"].Val)
	fmt.Printf("Decoded Result: %v\n", obj1x)
	fmt.Println("US==", obj1x[USD]) // 'US'

	decode(&obj2x, DataBase["list2"].Val)
	fmt.Printf("Decoded Result obj2: %#v\n", obj2x)
	fmt.Printf("Decoded Result obj2[0]: %#v\n", obj2x[0])

	// dicst
	dict1x := map[string]string{}

	decode(&dict1x, DataBase["dict1"].Val)
	fmt.Printf("Decoded Result dict1x: %#v\n", dict1x)
	fmt.Printf("Decoded Item of dict1x: %#v\n", dict1x["key1"])

}
