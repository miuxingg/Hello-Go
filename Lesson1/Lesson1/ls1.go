package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Hello world")

	// ARRAY
	fmt.Println("\n---------- ARRAY")
	intArray := [5]int{1, 2, 3, 4, 5}

	for i := 0; i < len(intArray); i++ {
		fmt.Print(intArray[i])
	}
	fmt.Println("\n----------")
	for index, element := range intArray {
		fmt.Printf("%v: %v \t", index, element)
	}
	fmt.Println("\n----------")
	for _, element := range intArray {
		fmt.Print(element)

	}

	// Copy array
	fmt.Println("\n---------- COPY ARRAY")
	arr1 := [3]string{"VN", "EN", "DE"}
	arr2 := arr1
	arr3 := &arr1
	arr1[1] = "CN"

	fmt.Printf("Arr1: %v\n", arr1)
	fmt.Printf("Arr2: %v\n", arr2)
	fmt.Printf("Arr3: %v\n", *arr3)
	fmt.Printf("Arr2 [1: 2]: %v", arr2[1:2])

	// SLIDE
	fmt.Println("\n---------- SLIDE")
	var intSlideVar []int
	intSlideShort := make([]int, 3, 5)
	fmt.Printf("intSlideVar %v\n", reflect.ValueOf(intSlideVar).Kind())
	fmt.Printf("intSlideShort %v\n", reflect.ValueOf(intSlideShort).Kind())
	fmt.Printf("intSlide %v, len(%v), cap(%v)\n", intSlideShort, len(intSlideShort), cap(intSlideShort))

	// Append
	slideAppend := append(intSlideShort, 10, 20, 30)
	fmt.Printf("intSlide append%v, len(%v), cap(%v)\n", slideAppend, len(slideAppend), cap(slideAppend))
	fmt.Printf("intSlide %v, len(%v), cap(%v)\n", intSlideShort, len(intSlideShort), cap(intSlideShort))

	// Remove
	slideRemove := append(intSlideShort[:1], intSlideShort[2])
	fmt.Printf("intSlide remove%v, len(%v), cap(%v)\n", slideRemove, len(slideRemove), cap(slideRemove))
	fmt.Printf("intSlide %v, len(%v), cap(%v)\n", intSlideShort, len(intSlideShort), cap(intSlideShort))

	// Append slide to slide
	intSlideShort = append(intSlideShort, intSlideVar...)
	fmt.Printf("intSlide double append %v, len(%v), cap(%v)\n", intSlideShort, len(intSlideShort), cap(intSlideShort))

	// MAP
	fmt.Println("\n---------- MAP")
	var map1 = map[string]int{}
	map2 := make(map[string]string)
	map1["x"] = 1
	map1["y"] = 2
	map2["x"] = "x"
	map2["y"] = "y"

	fmt.Printf("Map 1: %v , Map 2: %v\n", map1, map2)
	delete(map1, "x")
	fmt.Printf("Map 1 after delete: %v \n", map1)
	fmt.Printf("Map 2 loop: \n")
	for key, value := range map2 {
		fmt.Printf("key: %v, value: %v\t", key, value)
	}

}
