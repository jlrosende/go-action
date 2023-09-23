package utils

import "strings"

func Contains(s []string, str string) bool {
	for _, v := range s {
		if strings.EqualFold(v, str) {
			return true
		}
	}

	return false
}

func RemoveDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func ParseRuntime(runtime string) string {
	switch strings.ToLower(runtime) {
	case "java11", "java17", "graalvm11", "graalvm17":
		return "java"
	case "node16", "node18":
		return "node"
	case "go119", "go120":
		return "go"
	default:
		return "java"
	}
}
