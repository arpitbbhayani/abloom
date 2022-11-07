package main

import (
	"log"

	"github.com/arpitbbhayani/abloom"
)

func main() {
	b := abloom.NewBloom(512)
	b.Put([]byte("apple"))
	b.Put([]byte("banana"))
	b.Put([]byte("cat"))

	v, err := b.Check([]byte("apple"))
	if err != nil {
		log.Fatal("error while computing the hash")
	}
	log.Println("is apple present?", v)
}
