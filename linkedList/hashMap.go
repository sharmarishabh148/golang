package main

import "fmt"

var BucketSize int = 89 // prime number

//https://leetcode.com/problems/design-hashmap/solutions/1310141/golang-concise-solution/?languageTags=golang
type Data struct {
	Key   int
	Value int
}

type MyHashMap struct {
	Bucket [][]*Data
}

func Constructor() MyHashMap {
	b := make([][]*Data, BucketSize)
	return MyHashMap{
		Bucket: b,
	}
}

func (this *MyHashMap) Put(key int, value int) {
	// new
	bIdx := key % BucketSize
	if idx := this.index(key); idx == -1 {
		this.Bucket[bIdx] = append(this.Bucket[bIdx], &Data{key, value})
	} else {
		if this.Bucket[bIdx][idx] == nil {
			this.Bucket[bIdx] = append(this.Bucket[bIdx], &Data{key, value})
		} else {
			this.Bucket[bIdx][idx].Value = value
		}
	}
}

func (this *MyHashMap) Get(key int) int {
	if idx := this.index(key); idx != -1 {
		bIdx := key % BucketSize
		if this.Bucket[bIdx][idx] != nil {
			return this.Bucket[bIdx][idx].Value
		}
	}

	return -1
}

func (this *MyHashMap) Remove(key int) {
	if idx := this.index(key); idx != -1 {
		bIdx := key % BucketSize
		this.Bucket[bIdx][idx] = nil
	}
}

func (this *MyHashMap) index(key int) int {
	bIdx := key % BucketSize
	b := this.Bucket[bIdx]
	for i, d := range b {
		if d != nil && d.Key == key {
			return i
		}
	}
	return -1
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */

func main() {
	obj := MyHashMap{
		Bucket: make([][]*Data, BucketSize),
	}

	obj.Put(1, 1)
	obj.Put(2, 2)
	param_2 := obj.Get(1)
	fmt.Printf("param_2 : %v", param_2)
	param_2 = obj.Get(3)
	fmt.Printf("param_2 : %v", param_2)
	obj.Put(2, 1)
	param_2 = obj.Get(2)
	fmt.Printf("param_2 : %v", param_2)
	obj.Remove(2)
	param_2 = obj.Get(2)
	fmt.Printf("param_2 : %v", param_2)

}
