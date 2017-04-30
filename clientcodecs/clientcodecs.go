// kvdb_clientcodecs: encode/decode to process structures to []byte in the both ways
// part of KVDB database. copyright(C) GeneSemerenko

package clientcodecs

import (
	"bytes"
	"encoding/base64"
	"encoding/gob" // https://golang.org/pkg/encoding/gob/
)

// get interface structure end encode it to []byte result
// input: interface{}
// output: []byte encoded as base64

func Encode(obj interface{}) []byte {
	//var buf bytes.Buffer
	buf := new(bytes.Buffer)
	err := gob.NewEncoder(buf).Encode(obj)
	if err != nil {
		panic(err)
	}
	//return buf.Bytes() //ret []byte, old variant
	return []byte(base64.URLEncoding.EncodeToString(buf.Bytes()))
}

// get []byte source and interface structure end decode to interface{}
// output:  -> interface{}
// input: []byte

func Decode(objx interface{}, src []byte) {

	uDec, err := base64.URLEncoding.DecodeString(string(src))
	if err != nil {
		panic(err)
	}
	//fmt.Printf("got decoded:`%v`\n", uDec)
	// make a reader for the input (which is a []byte)
	p := bytes.NewReader(uDec)
	err = gob.NewDecoder(p).Decode(objx)

	if err != nil {
		panic(err)
	}
}
