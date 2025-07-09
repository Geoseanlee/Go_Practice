package main

func getNameCounts(names []string) map[rune]map[string]int {

	res := make(map[rune]map[string]int)

	for _, name := range names {
		if len(name) == 0 {
			continue
		}

		firstChar := []rune(name)[0] // []rune(name)：将字符串 name 转换为一个 rune 切片（[]rune）

		if res[firstChar] == nil {
			res[firstChar] = make(map[string]int)

		}

		res[firstChar][name]++
	}
	return res
}
