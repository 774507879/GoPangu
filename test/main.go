package main

import "diskdb"

func main() {
	db, _ := diskdb.Open("D:/db/panguDb")
	//db.Put([]byte("pangu"),[]byte("kaitian"))
	//db.Put([]byte("pzx"),[]byte("niubi"))
	db.Get([]byte("pzx"))
}
