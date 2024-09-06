package array

import (
	"strconv"
	"strings"
)

type ArrayTool struct{}

// 合并数组
func (s *ArrayTool) MergeArray(dest []interface{}, src []interface{}) (result []interface{}) {
	result = make([]interface{}, len(dest)+len(src))
	copy(result, dest)
	copy(result[len(dest):], src)
	return
}

// 删除数组
func (s *ArrayTool) DeleteArray(src []interface{}, index int) (result []interface{}) {
	result = append(src[:index], src[(index+1):]...)
	return
}

// []string => []int
func (s *ArrayTool) ArrayStrToInt(data []string) []int {
	var (
		arr = make([]int, 0, len(data))
	)
	if len(data) == 0 {
		return arr
	}
	for i, _ := range data {
		var num, _ = strconv.Atoi(data[i])
		arr = append(arr, num)
	}
	return arr
}

// []int => []string
func (s *ArrayTool) ArrayIntToStr(data []int) []string {
	var (
		arr = make([]string, 0, len(data))
	)
	if len(data) == 0 {
		return arr
	}
	for i, _ := range data {
		arr = append(arr, strconv.Itoa(data[i]))
	}
	return arr
}

// str[TrimSpace] in string list
func (s *ArrayTool) TrimSpaceStrInArray(str string, data []string) bool {
	if len(data) > 0 {
		for _, row := range data {
			if str == strings.TrimSpace(row) {
				return true
			}
		}
	}
	return false
}

// str in string list
func (s *ArrayTool) StrInArray(str string, data []string) bool {
	if len(data) > 0 {
		for _, row := range data {
			if str == row {
				return true
			}
		}
	}
	return false
}

// str in int list
func (s *ArrayTool) IntInArray(num int, data []int) bool {
	if len(data) > 0 {
		for _, row := range data {
			if num == row {
				return true
			}
		}
	}
	return false
}

/**
 * 数组（字符串）判断是否包含
 */
func (s *ArrayTool) IsInArrayString(list []string, ele string) bool {
	for _, element := range list {
		if ele == element {
			return true
		}
	}
	return false
}

/**
 * 数组（字符串）合并
 */
func (s *ArrayTool) ArrayStringMerge(a, b []string) []string {
	var arr []string
	for _, i := range a {
		arr = append(arr, i)
	}
	for _, j := range b {
		arr = append(arr, j)
	}
	return arr
}

/**
 * 数组（字符串）去重
 */
func (s *ArrayTool) ArrayStringUnique(m []string) []string {
	d := make([]string, 0)
	tempMap := make(map[string]bool, len(m))
	for _, v := range m {
		if !tempMap[v] {
			tempMap[v] = true
			d = append(d, v)
		}
	}
	return d
}

func (s *ArrayTool) StrArrayRemoveRepeatElement(list []string) []string {
	// 创建一个临时map用来存储数组元素
	temp := make(map[string]struct{})
	index := 0
	// 将元素放入map中
	for _, v := range list {
		temp[v] = struct{}{}
	}
	tempList := make([]string, len(temp))
	for key := range temp {
		tempList[index] = key
		index++
	}
	return tempList
}

/**
 * 数组（泛型）去重
 */
func RemoveRepeatElement[T comparable](list []T) []T {
	// 创建一个临时map用来存储数组元素
	temp := make(map[T]struct{})
	index := 0
	// 将元素放入map中
	for _, v := range list {
		temp[v] = struct{}{}
	}
	tempList := make([]T, len(temp))
	for key := range temp {
		tempList[index] = key
		index++
	}
	return tempList
}

// list filter out data that does not meet the criteria
func (s *ArrayTool) Filter(list []interface{}, condition func(interface{}) bool) []interface{} {
	filteredList := make([]interface{}, 0)
	for _, item := range list {
		if condition(item) {
			filteredList = append(filteredList, item)
		}
	}
	return filteredList
}
