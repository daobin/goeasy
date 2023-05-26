package internal

import (
	"path"
	"strings"
)

func MergeString(firstString string, otherStrings ...string) string {
	if otherStrings == nil || len(otherStrings) == 0 {
		return firstString
	}

	firstBytes := []byte(firstString)

	newBytes := make([]byte, len(firstBytes))
	copy(newBytes, firstBytes)

	for _, otherStr := range otherStrings {
		newBytes = append(newBytes, []byte(otherStr)...)
	}

	return string(newBytes)
}

func JoinPath(firstPath, secondPath string) string {
	firstPath = strings.Trim(firstPath, "/")
	secondPath = strings.Trim(secondPath, "/")
	if secondPath == "" {
		return MergeString("/", firstPath)
	}

	finalPath := path.Join(firstPath, secondPath)
	return MergeString("/", finalPath)
}

func FindPathWildcard(pathStr string) (string, int, bool) {
	for start, char := range []byte(pathStr) {
		if char != ':' && char != '*' {
			continue
		}

		valid := true
		for end, char := range []byte(pathStr[start+1:]) {
			switch char {
			case '/':
				return pathStr[start : start+1+end], start, valid
			case ':', '*':
				valid = false
			}
		}

		return pathStr[start:], start, valid
	}

	return "", -1, false
}
