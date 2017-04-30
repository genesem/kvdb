// Codes: aka encode/decode to process structures to []byte in the both ways

package main

import (
	"bytes"
	"encoding/gob" // https://golang.org/pkg/encoding/gob/
)

// get interface structure end encode it to []byte result
// input: interface{}
// output: []byte

func encode(obj interface{}) []byte {
	//var buf bytes.Buffer
	buf := new(bytes.Buffer)
	err := gob.NewEncoder(buf).Encode(obj)
	if err != nil {
		panic(err)
	}
	return buf.Bytes() //get the result out as a []byte
}

// get []byte source and interface structure end decode to interface{}
// output: interface{}
// input: []byte

func decode(objx interface{}, src []byte) {
	// make a reader for the input (which is a []byte)
	p := bytes.NewReader(src)
	err := gob.NewDecoder(p).Decode(objx)

	if err != nil {
		panic(err)
	}
}

// todo: panic recover @decode error, must send back errror and continue
