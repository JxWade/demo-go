package main

import (
	"demo-go/cache"
	"demo-go/poker"
	"demo-go/tools"
	"time"

	"encoding/base64"
	"fmt"
)

func main() {
	// 1. 实现一个函数，输入为任意长度的int64数组，返回元素最大差值，例如输入arr=[5,8,10,1,3]，返回9。
	fmt.Println("1. 实现一个函数，输入为任意长度的int64数组，返回元素最大差值，例如输入arr=[5,8,10,1,3]，返回9。")
	arr := tools.CreateRandInt64Slice()
	demo1 := tools.MaxDifference(arr)
	fmt.Println(demo1)
	fmt.Println("-----------------------------------------------------------------------------------")

	// 2. 实现一个函数，对输入的扑克牌执行洗牌，保证其是均匀分布的，也就是说列表中的每一张扑克牌出现在列表的每一个位置上的概率必须相同。
	fmt.Println("2. 实现一个函数，对输入的扑克牌执行洗牌，保证其是均匀分布的，也就是说列表中的每一张扑克牌出现在列表的每一个位置上的概率必须相同。")
	//初始化poker牌
	pokers := poker.NewPokers()
	pokers.Shuffle() //打乱

	//截取一段poker牌
	demo2 := poker.CustomShuffle(pokers.Pokers[:9])
	poker.EachOtherPokersTest(demo2)
	fmt.Println("-----------------------------------------------------------------------------------")

	// 3. 设计一个带失效时间的缓存数据结构，key和value都是string，并实现增删改查接口。
	fmt.Println("3. 设计一个带失效时间的缓存数据结构，key和value都是string，并实现增删改查接口。")
	cache.InitCache()
	store := cache.GetCacheStore()
	store.SetString("key1", "aaa", 2)
	key1 := store.GetString("key1")
	fmt.Println("key1 val:", *key1)
	store.SetString("key1", "bbb", 5)
	key11 := store.GetString("key1")
	fmt.Println("key1 reset val:", *key11)
	store.SetString("key2", "ccc", 0)
	key2 := store.GetString("key2")
	fmt.Println("key2 val:", *key2)
	kd1 := store.Get("key1")
	kd2 := store.Get("key2")
	fmt.Println("key1 val data:", *kd1)
	fmt.Println("key2 val data:", *kd2)
	time.Sleep(6 * time.Second)
	skd1 := store.Get("key1")
	skd2 := store.Get("key2")
	fmt.Println("sleep key1 val data:", skd1)
	fmt.Println("sleep key2 val data:", *skd2)
	store.RmString("key2")
	dkd2 := store.Get("key2")
	fmt.Println("remove key2 val data:", dkd2)
	fmt.Println("-----------------------------------------------------------------------------------")

	// 4. 实现一个游戏算法：输入n和m，代表有n个选手围成一圈（选手编号为0至n-1），0号从1开始报数，报m的选手游戏失败从圆圈中退出，下一个人接着从1开始报数，如此反复，求最后的胜利者编号。
	// 例如，n=3，m=2，那么失败者编号依次是1、0，最后的胜利者是2号。
	// 这里考虑m，n都是正常的数据范围，其中：
	// 1 <= n <= 10^5
	// 1 <= m <= 10^6
	//算法要求考虑时间效率。
	fmt.Println("4. 实现一个游戏算法：输入n和m，代表有n个选手围成一圈（选手编号为0至n-1），0号从1开始报数，报m的选手游戏失败从圆圈中退出")
	demo4 := tools.Game(10, 15)
	fmt.Println(demo4)
	fmt.Println("-----------------------------------------------------------------------------------")

	//如无不便，请优先使用github.com，gitlab.com或gitee.com提交答案，回复公开代码库地址即可。否则请使用zip压缩包形式提交代码。

	//请将答题发送至，邮箱(base64)：a29uZ2ppYW5AbGluZ2ppbmczNy5jb20=
	decoded, err := base64.StdEncoding.DecodeString("a29uZ2ppYW5AbGluZ2ppbmczNy5jb20=")
	if err != nil {
		fmt.Println("decode error:", err)
		return
	}
	fmt.Println(string(decoded))
	// kongjian@lingjing37.com
}
