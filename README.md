KVDB [![Build Status](https://travis-ci.org/genesem/kvdb.svg?branch=master)](https://travis-ci.org/genesem/kvdb)
========

KeyValue Nano Database written in Go (golang).

#### Storage types:

* Strings
* Lists of Strings
* Dictionaries of Strings

#### Features:

* Each key has TTL (time to live) is seconds and autodeleted when TTL got expired.
* Low RAM footprint.
* Fast work.

#### Install:

  `git clone https://github.com/genesem/kvdb`

  `cd kvdb/server && go build -o ../server`
  
  `runnable server file is: 'server'`

#### Usage:

By default server listens on :3000 tcp port.



#### Notes:

to be released later.