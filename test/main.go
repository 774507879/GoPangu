package main

import "diskdb"

func main() {
	db, _ := diskdb.Open("D:/db/panguDb")
	//db.Put([]byte("hello"),[]byte("world"))
	//db.Put([]byte("你好"),[]byte("世界"))
	db.Get([]byte("hello"))
}
