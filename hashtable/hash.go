package main

import (
	"fmt"
)

// simple CRC-like hash function
func hash(s string) uint {
	h := uint(0)
	for _, c := range s {
		v := uint(c)
		high := h & 0xf8000000
		h = h << 5
		h = h ^ (high >> 27)
		h = h ^ v
	}
	return h
}

func hashb(s string, buckets uint) uint {
	return hash(s) % buckets
}

type HashTable struct {
	size    uint
	values  uint
	buckets [][][2]string
}

func create(size ...int) *HashTable {
	h := &HashTable{}
	h.size = 127
	if len(size) > 0 {
		h.size = uint(size[0])
	}
	h.buckets = make([][][2]string, h.size)
	h.values = 0
	return h
}

// expand the size of the hash table, re-hashing based on the new bucketsize
func expand(h *HashTable) {
	old := h.buckets
	h.size = h.size*2 + 1
	h.values = 0
	h.buckets = make([][][2]string, h.size)
	for _, v := range old {
		for _, pair := range v {
			h.Insert(pair[0], pair[1])
		}
	}
}

func (h *HashTable) Delete(key string) {
	h.Pop(key)
}

func (h *HashTable) Has(key string) bool {
	_, err := h.GetErr(key)
	return err != nil
}

func (h *HashTable) Pop(key string) string {
	hashed := hashb(key, h.size)
	for i, v := range h.buckets[hashed] {
		if key == v[0] {
			ret := v[1]
			// remove from the bucket
			sub := h.buckets[hashed]
			sub[i] = sub[len(sub)-1]
			h.buckets[hashed] = sub[0 : len(sub)-1]
			h.values--
			return ret
		}
	}
	return ""
}

func (h *HashTable) GetErr(key string) (string, error) {
	hashed := hashb(key, h.size)
	sub := h.buckets[hashed]
	for _, v := range sub {
		if key == v[0] {
			return v[1], nil
		}
	}
	return "", fmt.Errorf("Not Found.")
}

func (h *HashTable) Get(key string) string {
	val, _ := h.GetErr(key)
	return val
}

func (h *HashTable) Insert(key, value string) {
	hashed := hashb(key, h.size)
	h.buckets[hashed] = append(h.buckets[hashed], [2]string{key, value})
	h.values++

	if h.values > h.size/2 {
		expand(h)
	}
}

func main() {

	tab := create()
	tab.Insert("Foo", "Foo is the word.")
	tab.Insert("Bar", "Bar is the word.")
	tab.Insert("Name", "Jason Moiron")
	tab.Insert("Greeting", "Hello, world!")

	fmt.Println(tab.Get("Foo"))
	fmt.Println(tab.Get("Bar"))
	fmt.Println(tab.Get("Name"))
	fmt.Println(tab.Get("DNE"))

	fmt.Println(tab.Pop("Name"))
	fmt.Println(tab.Pop("Name"))

}
