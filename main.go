package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {

	//查找只出现一次的数字
	go question_01()
	time.Sleep(time.Second)
	//有效括号
	go quest_02()
	time.Sleep(time.Second)
	//回文数
	go quest_03()
	time.Sleep(time.Second)
	//匹配字符相同字符最长的字符串
	go quest_04()
	time.Sleep(time.Second)
	//删除重复数据-处理方案与
	go question_01()
	time.Sleep(time.Second)
	//数组操作、进位处理
	quest_05()
	//删除有序数组中的重复项
	quest_06()
	//合并区间
	quest_07()
	//两数之和
	quest_08()

}

func quest_08() {
	nums := []int{1, 5, 6, 2, 9, 8, 10, 3, 2, 4, 0, 7}
	var target int = 19
	var flg []int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				flg = []int{nums[i], nums[j]}
				break
			}
		}
	}

	fmt.Println(flg)
}

func quest_07() {
	intervals := [][]int{{1, 5}, {4, 5}, {8, 9}, {15, 10}}
	var intervals_num [][]int
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })

	for k, v := range intervals {
		if k+1 >= len(intervals) {
			break
		}

		if v[1] >= intervals[k+1][0] {
			if v[0] < intervals[k+1][1] {

				pars := []int{v[0], intervals[k+1][1]}
				intervals_num = append(intervals_num, pars)
			}

		}

	}

	fmt.Println(intervals_num)
}

func quest_06() {
	num := []int{0, 1, 2, 4, 8, 2, 3, 5, 6, 4, 2, 1}

	l := len(num)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			//和下一个元素比较
			if num[i] == num[j] {
				//删除当前元素
				num = append(num[:i], num[i+1:]...)
				l-- //删除后元素长度更改
				i-- //删除后循环标识更改
				break
			}

		}
	}

	fmt.Println(num)
}

func quest_05() {

	arr01 := [...]int{2, 5, 6, 4, 1}
	arr_len := len(arr01)
	var result string
	if arr_len == 1 {
		result = strconv.Itoa(arr01[0] + 1)
	} else {
		for i := 0; i < len(arr01); i++ {
			var sss string
			if i == len(arr01)-1 {
				sss = strconv.Itoa(arr01[i] + 1)
			} else {
				sss = strconv.Itoa(arr01[i])
			}
			result = result + sss

		}
	}
	print := strings.Split(result, "")
	fmt.Println(print)
}

func quest_04() {
	strs := []string{"adsdge", "adsddeegy", "adsdfee", "adsefef", "adsssddeegyf"}
	sort.Slice(strs, func(i, j int) bool {
		return len(strs[i]) < len(strs[j])
	})

	var small string
	var srt_flg int
	for i := 0; i < len(strs[0]); i++ {
		small = small + string(strs[0][i])
		for j := 1; j < len(strs); j++ {
			index := strings.Index(strs[j], small)
			if index != -1 {
				srt_flg = i
			}
		}
	}

	fmt.Println("最长相同字符串是", strs[srt_flg])
}

func quest_03() {
	strs := "36636"
	lenth := int(len(strs) / 2)
	subStr1 := strs[0:lenth]
	subStr2 := strs[len(strs)-lenth:]
	if subStr1 == subStr2 {
		fmt.Println("是回文数")
	} else {
		fmt.Println("不是回文数")
	}
}

func quest_02() {
	var quest02 string = "(){>[]<>"
	var bs bool = true
	quest02_byte := []byte(quest02)

	for i := 0; i < len(quest02_byte); i++ {

		if i%2 == 0 {
			b := string(quest02_byte[i+1])
			var rg string
			switch quest02_byte[i] {

			case '(':
				rg = ")"
			case '[':
				rg = "]"
			case '<':
				rg = ">"
			case '{':
				rg = "}"
			}
			bs = quest02_test(b, rg)
			if !bs {
				bs = false
				break
			}

		}

	}
	fmt.Println(bs)
}

func quest02_test(a string, b string) bool {
	return a == b
}

func question_01() {

	quest01 := [...]int{1, 2, 2, 3, 8, 22, 5, 8, 6, 11, 22, 33, 44, 1}
	quest01_map := make(map[int]int)
	var result []int
	for i := 0; i < len(quest01); i++ {
		quest01_map[quest01[i]] = 1
		//fmt.Println(quest01[i])
		for s := 0; s < len(quest01); s++ {

			if i == s {
				continue
			}

			if quest01[i] == quest01[s] {

				quest01_map[quest01[i]] = quest01_map[quest01[i]] + 1
			}

		}
	}

	for k, v := range quest01_map {

		if v == 1 {
			result = append(result, k)
		}

	}

	fmt.Println(result)
}
