package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	fmt.Println(groupAnagramString([]string{"kita", "atik", "tika", "aku", "kia", "makan", "kua"}))
}

func groupAnagramString(arrayOfString []string) (res [][]string) {
	anagramMap := make(map[string][]string, 0)

	for idx, str := range arrayOfString {
		sortedStr := sortedSlice(str)
		if idx == 0 {
			anagramMap[sortedStr] = append(anagramMap[sortedStr], str)
			continue
		}
		if isAnagram(arrayOfString[idx-1], str) {
			anagramMap[sortedStr] = append(anagramMap[sortedStr], str)
			continue
		}
		anagramMap[sortedStr] = append(anagramMap[sortedStr], str)
	}

	for _, anagram := range anagramMap {
		res = append(res, anagram)
	}
	return res
}

func isAnagram(str string, str2 string) bool {
	if len(str) != len(str2) {
		return false
	}

	if sortedSlice(str) != sortedSlice(str2) {
		return false
	}

	return true
}

func sortedSlice(str string) string {
	strSlice := strings.Split(str, "")

	sort.Slice(strSlice, func(i, j int) bool {
		return strSlice[i] < strSlice[j]
	})

	return strings.Join(strSlice, "")
}

func isInList(str string, arrayOfString []string) bool {
	for _, strArr := range arrayOfString {
		if str == strArr {
			return true
		}
	}

	return false
}
