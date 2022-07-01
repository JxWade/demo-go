package tools

import (
	"math/rand"
	"sort"
	"time"
)

// Int64Slice 实现int64的排序
type Int64Slice []int64

func (x Int64Slice) Len() int           { return len(x) }
func (x Int64Slice) Less(i, j int) bool { return x[i] < x[j] }
func (x Int64Slice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

// MaxDifference
// 实现一个函数，输入为任意长度的int64数组，返回元素最大差值，例如输入arr=[5,8,10,1,3]，返回9。
func MaxDifference(arr []int64) (max int64) {
	// 数组排序
	sort.Sort(Int64Slice(arr))
	// 返回最大值-最小值
	return arr[len(arr)-1] - arr[0]
}

// CreateRandInt64Slice
// 生成随机数组
func CreateRandInt64Slice() []int64 {
	rand.Seed(time.Now().UnixNano())
	l := rand.Intn(100) //随机生成长度
	var arr []int64
	for i := 0; i < l; i++ {
		num := rand.Int63n(1000) //随机生成值
		arr = append(arr, num)
	}
	return arr
}
