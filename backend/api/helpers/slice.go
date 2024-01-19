package helpers

import "strings"

func KeyWordStr(sliceStr *[]string) *[]string {
	resultRemoveEmptyStrKeyword := RemoveEmptyStrKeyword(sliceStr)
	result := RemoveDuplicateStr(resultRemoveEmptyStrKeyword)
	return result
}

func RemoveEmptyStrKeyword(strSlice *[]string) *[]string {
	var list []string
	for _, item := range *strSlice {
		if item != "" && item != " " && item != "," && item != "." && item != "，" {
			item = strings.ReplaceAll(item, " ", "")
			list = append(list, item)
		}
	}

	return &list
}

func RemoveDuplicateStr(strSlice *[]string) *[]string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range *strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return &list
}
