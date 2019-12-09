package main

import "fmt"

/*
	马拉松算法，求最长回文子串，时间复杂度：线性
*/
func main() {
	// 回文数
	str := "abcddcbadcbadcabdadacd"

	// 填充#变成奇数个元素
	strArray := make([]byte, 0, 2*len(str)+1) // 每个字符是一个byte

	for i := 0; i < len(str); i++ {
		strArray = append(strArray, str[i])
		strArray = append(strArray, '#')
	}

	fmt.Print("原始字符串：")
	for i := 0; i < len(strArray); i++ {
		fmt.Print(string(strArray[i]))
	}
	fmt.Println()

	// 每个字符的最大回文半径
	radiusLen := make([]int, len(strArray))

	// 最大回文半径的中心位置
	id := 0

	// 最大回文串的右边界
	maxIndex := 0

	// 遍历新的串
	for i := 0; i < len(strArray); i++ {
		// 如果i在最大回文串中，那么可以进行判断，加快算法效率
		if i < maxIndex {
			j := 2*id - i // j和i是id的对称点

			if radiusLen[j] < maxIndex-i {
				// j的半径被最长串包住，那么i的半径必然等于j
				radiusLen[i] = radiusLen[j]
				continue
			} else if radiusLen[j] > maxIndex-i {
				// j的半径超出了最长串，那么i的半径必然是 j-(id-radiusLen(id)) = maxIndex - i 可画图观察
				radiusLen[i] = maxIndex - i
				continue
			} else if radiusLen[j] == maxIndex-i {
				// j的半径刚刚好到达最长串边界，这时i的半径可能比j还大，循环不会退出
				radiusLen[i] = radiusLen[j]
			}
		}

		for {
			// i半径必须合理，不能超过数组界，以圆心向两边拓展，逐一比较字符是否相等
			if i-radiusLen[i] >= 0 && i+radiusLen[i] < len(strArray) && strArray[i-radiusLen[i]] == strArray[i+radiusLen[i]] {
				radiusLen[i] = radiusLen[i] + 1
			} else {
				break
			}
		}

		// 如果半径比最大串还大，换人！
		if radiusLen[i] > radiusLen[id] {
			maxIndex = i + radiusLen[i] - 1
			id = i
		}
	}

	fmt.Print("处理完最长回文子串：")
	for i := id - (radiusLen[id] - 1); i <= id+(radiusLen[id]-1); i++ {
		fmt.Print(string(strArray[i]))
	}
}

