package lru

import "strings"

type SkipLRU struct{}

func (SkipLRU) LRUCache(calls []string) string {
	if len(calls) == 0 {
		return ""
	}

	reversedResult := make([]string, 0, cacheSize)

	// from the end of calls, if not in result - add
	for i := len(calls) - 1; i >= 0; i-- {
		call := calls[i]
		if !contains(reversedResult, call) {
			reversedResult = append(reversedResult, call)
		}

		// got full result - breaking
		if len(reversedResult) == cap(reversedResult) {
			break
		}
	}

	// reversing and writing
	sb := strings.Builder{}
	sb.Grow(len(reversedResult) * 2)

	for i := len(reversedResult) - 1; i >= 0; i-- {
		sb.WriteString(reversedResult[i])
		sb.WriteRune('-')
	}
	res := sb.String()

	// trimming the last '-' rune
	return res[:len(res)-1]
}

func contains(arr []string, elem string) bool {
	for _, v := range arr {
		if v == elem {
			return true
		}
	}

	return false
}
