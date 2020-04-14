//map介绍
/*
ages := make(map[string]int)//make创建一个map

//字面量创建并初始化
ages := map[string]int{
	    "alice":   31,
		"charlie": 34,
}

*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	sortMap()
	//is_equal := equalMap(map[string]int{"A": 0}, map[string]int{"B": 32})
	fmt.Printf("is_equal:%v", is_equal)
}

//map排序
func sortMap() {
	ages := map[string]int{
		"alice":   31,
		"charlie": 34,
		"atom":    34,
	}

	//b := &ages["alice"] cannot take the address of ages["alice"] 

	ages["alice"] = 3
	var names []string
	//因为ages已知，更明确的方式是：
	// names := make([]string, 0, len(ages))

	for name := range ages {
		names = append(names, name)
	}

	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}

	checkname := "bob"
	if _, ok := ages[checkname]; !ok {
		fmt.Println(checkname + " is not exist in ages")
	} else {
		fmt.Println(checkname + " exist in ages")
	}
}

//比较map
func equalMap(x, y map[string]int) bool {

	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}


