package internal

import (
	"path"
	"strings"
)

// MinNumber 返回两个数值中最小值
func MinNumber(first, second int) int {
	if first < second {
		return first
	}
	return second
}

// CommonPrefixLongest 返回相同前缀最大长度
func CommonPrefixLongest(first, second string) int {
	longest := 0
	max := MinNumber(len(first), len(second))
	for longest < max && first[longest] == second[longest] {
		longest++
	}
	return longest
}

// MergeString 拼接字符串
func MergeString(first string, others ...string) string {
	if others == nil || len(others) == 0 {
		return first
	}

	firstBytes := []byte(first)

	newBytes := make([]byte, len(firstBytes))
	copy(newBytes, firstBytes)

	for _, other := range others {
		newBytes = append(newBytes, []byte(other)...)
	}

	return string(newBytes)
}

// JoinPath 拼接路径
func JoinPath(firstPath, secondPath string) string {
	firstPath = strings.Trim(firstPath, "/")
	secondPath = strings.Trim(secondPath, "/")
	if secondPath == "" {
		return MergeString("/", firstPath)
	}

	finalPath := path.Join(firstPath, secondPath)
	return MergeString("/", finalPath)
}

// FindPathWildcard 查找路径中的通配符
func FindPathWildcard(pathStr string) (string, int, bool) {
	// 按字符循环查找匹配
	for start, char := range []byte(pathStr) {
		if char != ':' && char != '*' {
			continue
		}

		valid := true
		for end, char := range []byte(pathStr[start+1:]) {
			switch char {
			case '/':
				if len(pathStr[start:start+1+end]) == 1 {
					valid = false
				}
				return pathStr[start : start+1+end], start, valid
			case ':', '*':
				valid = false
			}
		}

		if len(pathStr[start:]) == 1 {
			valid = false
		}
		return pathStr[start:], start, valid
	}

	// 未找到返回 -1
	return "", -1, false
}
