package internal

import (
	"path"
	"strings"
)

func MinNumber(first, second int) int {
	if first < second {
		return first
	}
	return second
}

func CommonPrefixLongest(first, second string) int {
	longest := 0
	max := MinNumber(len(first), len(second))
	for longest < max && first[longest] == second[longest] {
		longest++
	}
	return longest
}

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

	return "", -1, false
}
