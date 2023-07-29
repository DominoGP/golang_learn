package main

import (
	"fmt"
	"sort"
)

func main() {
	basicSort()
	diyTypeSort()
}

func basicSort() {
	//对int类型的切片进行排序
	intSlice := []int{8, 3, 9, 7, 1}
	sort.Ints(intSlice)
	fmt.Println("Sorted int slice:", intSlice)

	//对float64类型的切片进行排序
	floatSlice := []float64{5.6, 1.3, 3.4, 2.2, 7.8}
	sort.Float64s(floatSlice)
	fmt.Println("Sorted float slice:", floatSlice)

	//对string类型的切片进行排序
	stringSlice := []string{"apple", "cat", "banana", "dog"}
	sort.Strings(stringSlice)
	fmt.Println("Sorted string slice:", stringSlice)

	//判断int类型的切片是否已经按升序排序
	fmt.Println("Is int slice sorted:", sort.IntsAreSorted(intSlice))

	//判断float64类型的切片是否已经按升序排序
	fmt.Println("Is float slice sorted:", sort.Float64sAreSorted(floatSlice))

	//判断string类型的切片是否已经按升序排序
	fmt.Println("Is string slice sorted:", sort.StringsAreSorted(stringSlice))
}

// 创建自定义类型，实现对应的三个排序接口
type Person struct {
	Name string
	Age  int
}

type ByAge []Person

func (a ByAge) Len() int      { return len(a) }
func (a ByAge) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

//只对Age排序
// func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

// 多重排序
func (a ByAge) Less(i, j int) bool {
	if a[i].Age == a[j].Age {
		return a[i].Name < a[j].Name
	}
	return a[i].Age < a[j].Age
}

func diyTypeSort() {
	persons := []Person{
		{"Tom", 23},
		{"Alice", 18},
		{"Bob", 27},
		{"Jane", 21},
		{"Zoe", 21},
	}

	// 实现排序
	sort.Sort(ByAge(persons))
	fmt.Println("\nSorted by age:", persons)

	// 实现排序反转
	sort.Sort(sort.Reverse(ByAge(persons)))
	fmt.Println("\nReverse:", ByAge(persons))
}
