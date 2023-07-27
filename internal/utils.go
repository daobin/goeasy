package internal

import (
	"path"
	"strings"
)

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

// JoinPath 拼接路径，统一返回小写
func JoinPath(firstPath, secondPath string) string {
	firstPath = strings.Trim(firstPath, "/")
	secondPath = strings.Trim(secondPath, "/")
	if secondPath == "" {
		return strings.ToLower(MergeString("/", firstPath))
	}

	finalPath := path.Join(firstPath, secondPath)
	return strings.ToLower(MergeString("/", finalPath))
}
