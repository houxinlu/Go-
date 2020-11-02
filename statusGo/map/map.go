package main

import "fmt"

func main() {
	m := map[string]string{
		"name":    "houxinlu",
		"course":  "golang",
		"site":    "home",
		"quality": "notbad",
	}
	m2 := make(map[string]int) //空map m2 == empty map
	var m3 map[string]int      // m3 == nil
	fmt.Println(m, m2, m3)

	fmt.Println("Traversing map")
	for k, v := range m {
		fmt.Println(k, v)
	} // map 输出的 k,v 是无序输出的，hashmap

	fmt.Println("Getting values")
	nameName := m["name"]
	fmt.Println(nameName)
	if nameName, ok := m["naame"]; ok {
		fmt.Println(nameName)
	} else {
		fmt.Println("key does not exist")
	}

	fmt.Println("Deleting values")
	name, ok := m["name"]
	fmt.Println(name, ok)

	delete(m, "name")
	name, ok = m["name"] //之前已经定义过了， 所以只能用 =，:=是声明并赋值  =是赋值(前提是要先声明)
	fmt.Println(name, ok)
}
// Map的操作
// map[k][v] map[k]map[k][v]
// 创建: make(map[string]int)
// 获取元素: m[key]
// key不存在时，获得Value类型的初始值
// 用value,ok := m[key] 来判断是否存在key
// 用delete删除一个key

//Map的遍历
//使用range遍历key,或者遍历key,value对
//不保证遍历顺序，如需顺序，需手动对key排序（加到list中，对list进行排序）
//使用len获得元素个数

//Map的key
//map使用哈希表，必须可以比较相等
//除了slice,map,function的内建类型都可以作为key
//Struct类型不包含上述字段，也可作为key



