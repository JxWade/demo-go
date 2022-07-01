package main

import (
	"demo-go/cache"
	"fmt"
)

func main() {
	//启动缓存系统
	cache.InitCache()
	store := cache.GetCacheStore()
	a := store.SetString("test", "aaa", 1000)
	fmt.Println(a)
	b := store.GetString("test")
	fmt.Println(*b)
	c := store.SetString("test", "bbb", 10000)
	fmt.Println(c)
	d := store.GetString("test")
	fmt.Println(*d)
	f := store.Get("test")
	fmt.Println(*f)
}

// Add
// 设计一个带失效时间的缓存数据结构，key和value都是string，并实现增删改查接口。
func Add(arr []int64) (max int64) {
	// 数组排序
	// 返回最大值-最小值
	return arr[len(arr)-1] - arr[0]
}
func Set(arr []int64) (max int64) {
	// 数组排序
	// 返回最大值-最小值
	return arr[len(arr)-1] - arr[0]
}
func Get(arr []int64) (max int64) {
	// 数组排序
	// 返回最大值-最小值
	return arr[len(arr)-1] - arr[0]
}

func Del(arr []int64) (max int64) {
	// 数组排序
	// 返回最大值-最小值
	return arr[len(arr)-1] - arr[0]
}
