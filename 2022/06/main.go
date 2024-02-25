package main

func find(input string) (index int) {
	found := false
	start := 0
	index = 4
	l := len(input)
	for !found {
		if start+index > l {
			break
		}
		m := make(map[rune]bool, 4)
		found = true
		s := input[start:index]
		for _, r := range s {
			if m[r] {
				found = false
				break
			}
			m[r] = true
		}
		if !found {
			start++
			index++
		}
	}
	return index
}
