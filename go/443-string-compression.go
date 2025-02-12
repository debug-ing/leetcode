package main

import "strconv"

func Compress(chars []byte) int {
	n := len(chars)
	if n <= 1 {
		return n
	}
	track := 0
	curr := chars[0]
	cnt := 1
	for i := 1; i < n; i++ {
		if chars[i-1] != chars[i] {
			chars[track] = curr
			track++

			if cnt > 1 {
				s := strconv.Itoa(cnt)
				for j := 0; j < len(s); j++ {
					chars[track] = s[j]
					track++
				}
			}
			curr = chars[i]
			cnt = 1
		} else {
			cnt++
		}
	}

	chars[track] = curr
	track++
	if cnt > 1 {
		s := strconv.Itoa(cnt)
		for j := 0; j < len(s); j++ {
			chars[track] = s[j]
			track++
		}
	}
	return track
}
