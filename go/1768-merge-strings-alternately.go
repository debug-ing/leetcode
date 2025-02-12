package main

func MergeAlternately(word1 string, word2 string) string {
	result := make([]byte, 0, len(word1)+len(word2))
	minLen := len(word1)
	if len(word2) < minLen {
		minLen = len(word2)
	}

	for i := 0; i < minLen; i++ {
		result = append(result, word1[i], word2[i])
	}
	result = append(result, []byte(word1[minLen:])...)
	result = append(result, []byte(word2[minLen:])...)
	return string(result)
}
