package main

import (
	"fmt"
	"sync"
)

type singleton map[string] string

var (
	once sync.Once
	instance singleton
	arr []int= []int{1,2,3}
)

func New() (singleton,*[]int){
	once.Do(
		func() {
			instance = make(singleton)
			//arr = []int{1,2}
		})
	return instance,&arr
}

func main() {

	s1,arr1 := New()
	s2,arr2 := New()
	s1["abc"] ="abc"


	fmt.Println(s2["abc"])
	fmt.Printf("%p\n",arr1)
	fmt.Printf("%p",arr2)
}
