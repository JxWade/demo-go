package tools

// Game
// 实现一个游戏算法：输入n和m，代表有n个选手围成一圈（选手编号为0至n-1），0号从1开始报数，报m的选手游戏失败从圆圈中退出，下一个人接着从1开始报数，如此反复，求最后的胜利者编号。
//例如，n=3，m=2，那么失败者编号依次是1、0，最后的胜利者是2号。
//这里考虑m，n都是正常的数据范围，其中：
//1 <= n <= 10^5
//1 <= m <= 10^6
//算法要求考虑时间效率。
func Game(n, m uint) (win uint) {
	// 初始化一个数组
	arr := make([]uint, n)
	for i := range arr {
		arr[i] = uint(i)
	}

	j := 0
	for {
		if len(arr) == 1 {
			break
		}
		i := uint(1)
		for {
			if j == len(arr) {
				j = 0
			}
			if i == m {
				arr = append(arr[:j], arr[j+1:]...)
				break
			}
			i++
			j++
		}
	}
	win = arr[0]
	return win
}
